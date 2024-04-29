package plugins

import (
	"custom-go/pkg/types"
	"custom-go/pkg/utils"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type httpProxyHookFunction func(*types.HttpTransportHookRequest, *HttpTransportBody) (*types.WunderGraphResponse, error)

func RegisterProxyHook(hookFunc httpProxyHookFunction, operationType ...types.OperationType) {
	callerName := utils.GetCallerName(string(types.HookParent_proxy))
	apiPath := strings.ReplaceAll(string(types.Endpoint_proxy), "{path}", callerName)

	types.AddEchoRouterFunc(func(e *echo.Echo) {
		e.Logger.Debugf(`Registered hookFunction [%s]`, apiPath)
		e.POST(apiPath, buildProxyFunc(hookFunc))
	})

	types.AddHealthFunc(func(e *echo.Echo, report *types.HealthReportLock) {
		operation := &types.Operation{}
		operationJsonPath := filepath.Join(string(types.HookParent_proxy), callerName) + jsonExtension

		// 读文件，保留原有配置，只需更新schema
		if !utils.NotExistFile(operationJsonPath) {
			_ = utils.ReadStructAndCacheFile(operationJsonPath, operation)
		} else {
			operation.Name = callerName
			operation.Path = apiPath
			operation.OperationType = types.OperationType_MUTATION
		}

		if operationType != nil && len(operationType) > 0 {
			operation.OperationType = operationType[0]
		}

		operationBytes, err := json.Marshal(operation)
		if err != nil {
			e.Logger.Errorf("json marshal failed, err: %v", err.Error())
			return
		}
		err = os.WriteFile(operationJsonPath, operationBytes, 0644)
		if err != nil {
			e.Logger.Errorf("write file failed, err: %v", err.Error())
			return
		}

		report.Lock()
		defer report.Unlock()
		report.Proxys = append(report.Proxys, callerName)
	})
}

func buildProxyFunc(proxyHook httpProxyHookFunction) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		brc := c.(*types.HttpTransportHookRequest)

		var reqBody HttpTransportBody
		err = c.Bind(&reqBody)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		newResp, err := proxyHook(brc, &reqBody)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		resp := types.MiddlewareHookResponse{
			Op:   reqBody.Name,
			Hook: types.MiddlewareHook(types.HookParent_proxy),
		}
		if newResp != nil {
			resp.Response = types.OnResponseHookResponse{Response: newResp}
		}
		return c.JSON(http.StatusOK, resp)
	}
}
