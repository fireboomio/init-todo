package plugins

import (
	"custom-go/pkg/types"
	"custom-go/pkg/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

const maximumRecursionLimit = 16

func ConvertBodyFunc[I, O any](oldFunc func(*types.HookRequest, *types.OperationBody[I, O]) (*types.OperationBody[I, O], error)) types.OperationHookFunction {
	return func(hook *types.HookRequest, body *types.OperationBody[any, any]) (res *types.OperationBody[any, any], err error) {
		// 将传入的 OperationBody 转换为需要的类型
		var input = utils.ConvertType[types.OperationBody[any, any], types.OperationBody[I, O]](body)
		// 调用旧函数获取结果
		oldRes, err := oldFunc(hook, input)
		if err != nil {
			return res, err
		}

		res = utils.ConvertType[types.OperationBody[I, O], types.OperationBody[any, any]](oldRes)
		return res, nil
	}
}

func RegisterOperationsHooks(e *echo.Echo, operations []string, operationHooksMap types.OperationHooks) {
	if len(operationHooksMap) == 0 {
		return
	}
	for _, operationPath := range operations {
		registerOperationHooks(e, operationPath, operationHooksMap)
	}
}

func MakeDataAnyMap(data any) map[string]any {
	return map[string]any{"data": data}
}

func registerOperationHooks(e *echo.Echo, operationPath string, operationHooksMap types.OperationHooks) {
	if operationHook, ok := operationHooksMap[operationPath]; ok {
		if operationHook.MockResolve != nil {
			apiPath := strings.ReplaceAll(string(types.Endpoint_mockResolve), "{path}", operationPath)
			e.Logger.Debugf(`Registered operationHook [%s]`, apiPath)
			e.POST(apiPath, buildOperationHook(operationPath, types.MiddlewareHook_mockResolve, operationHook.MockResolve, mockResolve))
		}

		if operationHook.PreResolve != nil {
			apiPath := strings.ReplaceAll(string(types.Endpoint_preResolve), "{path}", operationPath)
			e.Logger.Debugf(`Registered operationHook [%s]`, apiPath)
			e.POST(apiPath, buildOperationHook(operationPath, types.MiddlewareHook_preResolve, operationHook.PreResolve, preResolve))
		}

		if operationHook.PostResolve != nil {
			apiPath := strings.ReplaceAll(string(types.Endpoint_postResolve), "{path}", operationPath)
			e.Logger.Debugf(`Registered operationHook [%s]`, apiPath)
			e.POST(apiPath, buildOperationHook(operationPath, types.MiddlewareHook_postResolve, operationHook.PostResolve, postResolve))
		}

		if operationHook.MutatingPreResolve != nil {
			apiPath := strings.ReplaceAll(string(types.Endpoint_mutatingPreResolve), "{path}", operationPath)
			e.Logger.Debugf(`Registered operationHook [%s]`, apiPath)
			e.POST(apiPath, buildOperationHook(operationPath, types.MiddlewareHook_mutatingPreResolve, operationHook.MutatingPreResolve, mutatingPreResolve))
		}

		if operationHook.MutatingPostResolve != nil {
			apiPath := strings.ReplaceAll(string(types.Endpoint_mutatingPostResolve), "{path}", operationPath)
			e.Logger.Debugf(`Registered operationHook [%s]`, apiPath)
			e.POST(apiPath, buildOperationHook(operationPath, types.MiddlewareHook_mutatingPostResolve, operationHook.MutatingPostResolve, mutatingPostResolve))
		}

		if operationHook.CustomResolve != nil {
			apiPath := strings.ReplaceAll(string(types.Endpoint_customResolve), "{path}", operationPath)
			e.Logger.Debugf(`Registered operationHook [%s]`, apiPath)
			e.POST(apiPath, buildOperationHook(operationPath, types.MiddlewareHook_customResolve, operationHook.CustomResolve, customResolve))
		}
	}
}

func requestContext(c echo.Context) (result *types.HookRequest, err error) {
	body := make(map[string]interface{})
	if err := c.Request().ParseForm(); err != nil {
		return result, err
	}

	result = c.(*types.HookRequest)
	for key, value := range c.Request().Form {
		body[key] = value[0]
	}
	if cycleCounter, ok := body["cycleCounter"].(int); ok {
		if cycleCounter > maximumRecursionLimit {
			return result, fmt.Errorf("maximum recursion limit reached (%d)", maximumRecursionLimit)
		}
		result.InternalClient = result.InternalClient.WithHeaders(types.RequestHeaders{"Wg-Cycle-Counter": strconv.Itoa(cycleCounter)})
	}
	return result, nil
}

func mockResolve(in, out *types.OperationBody[any, any]) {
	in.Response = out.Response
	in.SetClientRequestHeaders = out.SetClientRequestHeaders
}
func preResolve(in, out *types.OperationBody[any, any]) {
	in.SetClientRequestHeaders = out.SetClientRequestHeaders
}

func postResolve(in, out *types.OperationBody[any, any]) {
	in.SetClientRequestHeaders = out.SetClientRequestHeaders
}

func mutatingPreResolve(in, out *types.OperationBody[any, any]) {
	in.Input = out.Input
	in.SetClientRequestHeaders = out.SetClientRequestHeaders
}

func mutatingPostResolve(in, out *types.OperationBody[any, any]) {
	in.Response = out.Response
	in.SetClientRequestHeaders = out.SetClientRequestHeaders
	if in.Response != nil && in.Response.DataAny != nil {
		in.Response.Data = in.Response.DataAny
		in.Response.DataAny = nil
	}
}

func customResolve(in, out *types.OperationBody[any, any]) {
	in.Response = out.Response
	in.SetClientRequestHeaders = out.SetClientRequestHeaders
}

func buildOperationHook(operationName string, hookName types.MiddlewareHook, hookFunction types.OperationHookFunction, action func(in, out *types.OperationBody[any, any])) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusOK)

		var in types.OperationBody[any, any]
		err = c.Bind(&in)
		if err != nil {
			return
		}

		hookRequest, err := requestContext(c)
		if err != nil {
			return
		}

		in.Op = operationName
		in.Hook = hookName
		in.SetClientRequestHeaders = HeadersToObject(c.Request().Header)
		out, err := hookFunction(hookRequest, &in)
		if err != nil {
			return err
		}

		if out != nil {
			action(&in, out)
		}
		return c.JSON(http.StatusOK, &in)
	}
}

func HeadersToObject(headers http.Header) types.RequestHeaders {
	obj := make(types.RequestHeaders)
	for key, values := range headers {
		if len(values) > 0 {
			obj[key] = values[0]
		}
	}
	return obj
}
