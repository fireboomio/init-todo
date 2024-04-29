package plugins

import (
	"custom-go/pkg/types"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"net/http"
)

type GlobalConfiguration struct {
	HttpTransport HttpTransportHooks
	WsTransport   WsTransportHooks
}

type HttpTransportBody struct {
	Request  *types.WunderGraphRequest  `json:"request"`
	Response *types.WunderGraphResponse `json:"response"`
	Name     string                     `json:"operationName"`
	Type     string                     `json:"operationType"`
}

type WsTransportBody struct {
	DataSourceId string `json:"dataSourceId"`
}

type HttpTransportHooks struct {
	BeforeOriginRequest func(*types.HttpTransportHookRequest, *HttpTransportBody) (*types.WunderGraphRequest, error)
	AfterOriginResponse func(*types.HttpTransportHookRequest, *HttpTransportBody) (*types.WunderGraphResponse, error)
	OnOriginRequest     func(*types.HttpTransportHookRequest, *HttpTransportBody) (*types.WunderGraphRequest, error)
	OnOriginResponse    func(*types.HttpTransportHookRequest, *HttpTransportBody) (*types.WunderGraphResponse, error)
}

type WsTransportHooks struct {
	OnConnectionInit func(*types.WsTransportHookRequest, *WsTransportBody) (any, error)
}

func RegisterGlobalHooks(e *echo.Echo, globalHooks GlobalConfiguration) {
	if globalHooks.HttpTransport.BeforeOriginRequest != nil {
		apiPath := string(types.Endpoint_beforeOriginRequest)
		e.Logger.Debugf(`Registered globalHook [%s]`, apiPath)
		e.POST(apiPath, func(c echo.Context) error {
			brc := c.(*types.HttpTransportHookRequest)
			var reqBody HttpTransportBody
			err := c.Bind(&reqBody)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			newReq, err := globalHooks.HttpTransport.BeforeOriginRequest(brc, &reqBody)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			resp := types.MiddlewareHookResponse{
				Op:   reqBody.Name,
				Hook: types.MiddlewareHook_beforeOriginRequest,
			}
			if newReq != nil {
				resp.Response = types.OnRequestHookResponse{Request: newReq}
			}
			return c.JSON(http.StatusOK, resp)
		})
	}

	if globalHooks.HttpTransport.AfterOriginResponse != nil {
		apiPath := string(types.Endpoint_afterOriginResponse)
		e.Logger.Debugf(`Registered globalHook [%s]`, apiPath)
		e.POST(apiPath, func(c echo.Context) error {
			brc := c.(*types.HttpTransportHookRequest)
			var respBody HttpTransportBody
			err := c.Bind(&respBody)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			newResp, err := globalHooks.HttpTransport.AfterOriginResponse(brc, &respBody)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			resp := types.MiddlewareHookResponse{
				Op:   respBody.Name,
				Hook: types.MiddlewareHook_afterOriginResponse,
			}
			if newResp != nil {
				resp.Response = types.OnResponseHookResponse{Response: newResp}
			}
			return c.JSON(http.StatusOK, resp)
		})
	}

	if globalHooks.HttpTransport.OnOriginRequest != nil {
		apiPath := string(types.Endpoint_onOriginRequest)
		e.Logger.Debugf(`Registered globalHook [%s]`, apiPath)
		e.POST(apiPath, func(c echo.Context) error {
			brc := c.(*types.HttpTransportHookRequest)
			var reqBody HttpTransportBody
			err := c.Bind(&reqBody)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			newReq, err := globalHooks.HttpTransport.OnOriginRequest(brc, &reqBody)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			resp := types.MiddlewareHookResponse{
				Op:   reqBody.Name,
				Hook: types.MiddlewareHook_onOriginRequest,
			}
			if newReq != nil {
				resp.Response = types.OnRequestHookResponse{Request: newReq}
			}
			return c.JSON(http.StatusOK, resp)
		})
	}

	if globalHooks.HttpTransport.OnOriginResponse != nil {
		apiPath := string(types.Endpoint_onOriginResponse)
		e.Logger.Debugf(`Registered globalHook [%s]`, apiPath)
		e.POST(apiPath, func(c echo.Context) error {
			brc := c.(*types.HttpTransportHookRequest)
			var respBody HttpTransportBody
			err := c.Bind(&respBody)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			newResp, err := globalHooks.HttpTransport.OnOriginResponse(brc, &respBody)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			resp := types.MiddlewareHookResponse{
				Op:   respBody.Name,
				Hook: types.MiddlewareHook_onOriginResponse,
			}
			if newResp != nil {
				resp.Response = types.OnResponseHookResponse{Response: newResp}
			}
			return c.JSON(http.StatusOK, resp)
		})
	}

	if globalHooks.WsTransport.OnConnectionInit != nil {
		apiPath := string(types.Endpoint_onConnectionInit)
		e.Logger.Debugf(`Registered globalHook [%s]`, apiPath)
		e.POST(apiPath, func(c echo.Context) error {
			brc := c.(*types.WsTransportHookRequest)
			var reqBody WsTransportBody
			err := c.Bind(&reqBody)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			resp, err := globalHooks.WsTransport.OnConnectionInit(brc, &reqBody)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return c.JSON(http.StatusOK, types.MiddlewareHookResponse{
				Hook:     types.MiddlewareHook_onConnectionInit,
				Response: resp,
			})
		})
	}

	// handle not found routes
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		var he *echo.HTTPError
		if errors.As(err, &he) {
			_ = c.JSON(he.Code, map[string]string{"error": cast.ToString(he.Message)})
		} else {
			_ = c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}
