package plugins

import (
	"custom-go/pkg/types"
	"custom-go/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type UploadHooks = map[string]UploadHooksProfile

type UploadBody[M any] struct {
	File  types.HookFile `json:"file"`
	Meta  M              `json:"meta"`
	Error struct {
		Name    string `json:"name"`
		Message string `json:"message"`
	} `json:"error"`
}

type UploadFunction func(request *types.UploadHookRequest, body *UploadBody[any]) (*types.UploadHookResponse, error)

type UploadHooksProfile struct {
	PreUpload  UploadFunction
	PostUpload UploadFunction
}

func ConvertUploadFunc[M any](oldFunc func(*types.UploadHookRequest, *UploadBody[M]) (*types.UploadHookResponse, error)) UploadFunction {
	return func(hook *types.UploadHookRequest, body *UploadBody[any]) (res *types.UploadHookResponse, err error) {
		// 将传入的 OperationBody 转换为需要的类型
		var input = utils.ConvertType[UploadBody[any], UploadBody[M]](body)
		// 调用旧函数获取结果
		return oldFunc(hook, input)
	}
}

func RegisterUploadsHooks(e *echo.Echo, uploadHooksMap map[string]UploadHooks) {
	for providerName, provider := range uploadHooksMap {
		for profileName, profile := range provider {
			if profile.PreUpload != nil {
				preUpload(e, providerName, profileName, profile.PreUpload)
			}
			if profile.PostUpload != nil {
				postUpload(e, providerName, profileName, profile.PostUpload)
			}
		}
	}
}

func preUpload(e *echo.Echo, providerName, profileName string, handler UploadFunction) {
	apiPath := strings.ReplaceAll(string(types.Endpoint_preUpload), "{provider}", providerName)
	apiPath = strings.ReplaceAll(apiPath, "{profile}", profileName)
	e.Logger.Debugf(`Registered uploadHook [%s]`, apiPath)
	e.POST(apiPath, func(c echo.Context) error {
		pur := c.(*types.UploadHookRequest)
		var param UploadBody[any]
		err := c.Bind(&param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}

		result, err := handler(pur, &param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, result)
	})
}

func postUpload(e *echo.Echo, providerName, profileName string, handler UploadFunction) {
	apiPath := strings.ReplaceAll(string(types.Endpoint_postUpload), "{provider}", providerName)
	apiPath = strings.ReplaceAll(apiPath, "{profile}", profileName)
	e.Logger.Debugf(`Registered uploadHook [%s]`, apiPath)
	e.POST(apiPath, func(c echo.Context) error {
		pur := c.(*types.UploadHookRequest)
		var param UploadBody[any]
		err := c.Bind(&param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}

		result, err := handler(pur, &param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, result)
	})
}
