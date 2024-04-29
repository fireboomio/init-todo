package types

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

var (
	PublicNodeUrl       string
	PrivateNodeUrl      string
	ServerListenAddress string
)

func (h *RequestHeaders) Get(key string) string {
	if h == nil {
		return ""
	}
	return (*h)[key]
}

func (r *WunderGraphRequest) NewRequest() *http.Request {
	req, _ := http.NewRequest(r.Method, r.RequestURI, bytes.NewReader(r.OriginBody))
	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}
	return req
}

func (r *WunderGraphResponse) Header() http.Header {
	return make(http.Header)
}

func (r *WunderGraphResponse) Write(i []byte) (int, error) {
	r.OriginBody = append(r.OriginBody, i...)
	return len(i), nil
}

func (r *WunderGraphResponse) WriteHeader(statusCode int) {
	r.StatusCode = int64(statusCode)
}

type (
	HealthReportLock struct {
		HealthReport
		sync.Mutex
	}
	BaseRequestContext struct {
		echo.Context
		*InternalClient
	}
	AuthenticationHookRequest = BaseRequestContext
	HookRequest               = BaseRequestContext
	HttpTransportHookRequest  = BaseRequestContext
	WsTransportHookRequest    = BaseRequestContext
	UploadHookRequest         = BaseRequestContext
)

type (
	registeredHook func(echo.Logger)
	healthFunc     func(*echo.Echo, *HealthReportLock)
	routerFunc     func(e *echo.Echo)
)

var (
	registeredHookArr []registeredHook
	healthFuncArr     []healthFunc
	routerFuncArr     []routerFunc
)

func GetRegisteredHookArr() []registeredHook {
	return registeredHookArr
}

func GetHealthFuncArr() []healthFunc {
	return healthFuncArr
}

func GetEchoRouterFuncArr() []routerFunc {
	return routerFuncArr
}

func AddRegisteredHook(hook registeredHook) {
	registeredHookArr = append(registeredHookArr, hook)
}

func AddHealthFunc(f healthFunc) {
	healthFuncArr = append(healthFuncArr, f)
}

func AddEchoRouterFunc(f routerFunc) {
	routerFuncArr = append(routerFuncArr, f)
}
