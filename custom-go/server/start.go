package server

import (
	"custom-go/pkg/plugins"
	"custom-go/pkg/types"
	"custom-go/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"context"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := startServer(); err != nil {
		os.Exit(1)
	}
}

func configureWunderGraphServer() *echo.Echo {
	// 初始化 Echo 实例
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	// 配置日志中间件
	loggerConfig := middleware.DefaultLoggerConfig
	loggerConfig.Skipper = func(c echo.Context) bool {
		return c.Request().URL.Path == "/health"
	}
	e.Use(middleware.LoggerWithConfig(loggerConfig))

	// 配置 CORS 中间件
	corsCfg := middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}
	e.Use(middleware.CORSWithConfig(corsCfg))

	plugins.RegisterGlobalHooks(e, plugins.WdgHooksAndServerConfig.Hooks.Global)
	plugins.RegisterAuthHooks(e, plugins.WdgHooksAndServerConfig.Hooks.Authentication)
	plugins.RegisterUploadsHooks(e, plugins.WdgHooksAndServerConfig.Hooks.Uploads)

	internalQueries := plugins.FetchOperations(e.Logger, types.OperationType_QUERY, true)
	if queryLen := len(internalQueries); queryLen > 0 {
		plugins.RegisterOperationsHooks(e, internalQueries, plugins.WdgHooksAndServerConfig.Hooks.Queries)
		e.Logger.Debugf(`Registered (%d) query operations`, queryLen)
	}

	internalMutations := plugins.FetchOperations(e.Logger, types.OperationType_MUTATION, true)
	if mutationLen := len(internalMutations); mutationLen > 0 {
		plugins.RegisterOperationsHooks(e, internalMutations, plugins.WdgHooksAndServerConfig.Hooks.Mutations)
		e.Logger.Debugf(`Registered (%d) mutation operations`, mutationLen)
	}

	subscriptionOperations := plugins.FetchOperations(e.Logger, types.OperationType_SUBSCRIPTION, false)
	if subscriptionLen := len(subscriptionOperations); subscriptionLen > 0 {
		plugins.RegisterOperationsHooks(e, subscriptionOperations, plugins.WdgHooksAndServerConfig.Hooks.Subscriptions)
		e.Logger.Debugf(`Registered (%d) subscription operations`, subscriptionLen)
	}

	registerOnce := &sync.Once{}
	e.Use(middleware.Recover(), func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			registerOnce.Do(func() {
				for _, registeredHook := range types.GetRegisteredHookArr() {
					go registeredHook(e.Logger)
				}
			})
			if c.Request().Method == http.MethodGet {
				return next(c)
			}

			var body types.BaseRequestBody
			err := utils.CopyAndBindRequestBody(c.Request(), &body)
			if err != nil {
				return err
			}

			if body.Wg == nil {
				body.Wg = &types.BaseRequestBodyWg{}
			}
			if body.Wg.ClientRequest == nil {
				body.Wg.ClientRequest = &types.WunderGraphRequest{
					Method:     c.Request().Method,
					RequestURI: c.Request().RequestURI,
					Headers:    plugins.HeadersToObject(c.Request().Header),
				}
			}
			headerRequestIdKey := string(types.InternalHeader_X_Request_Id)
			headerTraceIdKey := string(types.InternalHeader_X_uber_trace_id)
			internalClient := types.InternalClientFactoryCall(types.RequestHeaders{
				headerRequestIdKey: c.Request().Header.Get(headerRequestIdKey),
				headerTraceIdKey:   c.Request().Header.Get(headerTraceIdKey),
			}, body.Wg)
			brc := &types.BaseRequestContext{
				Context:        c,
				InternalClient: internalClient,
			}
			return next(brc)
		}
	})

	for _, routerFunc := range types.GetEchoRouterFuncArr() {
		routerFunc(e)
	}

	var healthReport *types.HealthReportLock
	e.Server.BaseContext = func(_ net.Listener) context.Context {
		healthReport = &types.HealthReportLock{}
		healthReport.Time = time.Now()
		for _, healthFunc := range types.GetHealthFuncArr() {
			go healthFunc(e, healthReport)
		}
		return context.Background()
	}
	workdir, _ := os.Getwd()
	// 健康检查
	e.GET(string(types.Endpoint_health), func(c echo.Context) error {
		return c.JSON(http.StatusOK, types.Health{
			Status:  "ok",
			Report:  &healthReport.HealthReport,
			Workdir: workdir,
		})
	})

	return e
}

func startServer() error {
	graphqlApi := types.WdgGraphConfig.Api
	types.PublicNodeUrl = types.GetConfigurationVal(graphqlApi.NodeOptions.PublicNodeUrl)
	types.PrivateNodeUrl = types.GetConfigurationVal(graphqlApi.NodeOptions.NodeUrl)
	serverListen := graphqlApi.ServerOptions.Listen
	types.ServerListenAddress = types.GetConfigurationVal(serverListen.Host) + ":" + types.GetConfigurationVal(serverListen.Port)

	// 配置服务器
	wdgServer := configureWunderGraphServer()

	// 启动服务器
	go func() {
		if err := wdgServer.Start(types.ServerListenAddress); err != nil {
			panic(err)
		}
	}()

	// 等待终止信号
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	// 优雅地关闭服务器
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := wdgServer.Shutdown(ctx); err != nil {
		panic(err)
	}

	return nil
}
