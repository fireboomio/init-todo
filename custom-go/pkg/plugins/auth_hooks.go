package plugins

import (
	"custom-go/pkg/types"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthenticationResponse struct {
	Message string      `json:"message"`
	Status  string      `json:"status"`
	User    *types.User `json:"user"`
}

type AuthenticationConfiguration struct {
	PostAuthentication         func(hook *types.AuthenticationHookRequest) error
	MutatingPostAuthentication func(hook *types.AuthenticationHookRequest) (*AuthenticationResponse, error)
	RevalidateAuthentication   func(hook *types.AuthenticationHookRequest) (*AuthenticationResponse, error)
	PostLogout                 func(hook *types.AuthenticationHookRequest) error
}

func RegisterAuthHooks(e *echo.Echo, authHooks AuthenticationConfiguration) {
	// preHandler hook - check user context
	e.Group("/authentication", func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			brc := c.(*types.AuthenticationHookRequest)
			if brc.User == nil {
				return errors.New("user context doesn't exist")
			}
			return next(brc)
		}
	})

	// authentication routes
	if authHooks.PostAuthentication != nil {
		apiPath := string(types.Endpoint_postAuthentication)
		e.Logger.Debugf(`Registered authHook [%s]`, apiPath)
		e.POST(apiPath, func(c echo.Context) error {
			brc := c.(*types.AuthenticationHookRequest)
			err := authHooks.PostAuthentication(brc)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			return c.JSON(http.StatusOK, types.MiddlewareHookResponse{
				Hook: types.MiddlewareHook_postAuthentication,
			})
		})
	}

	if authHooks.MutatingPostAuthentication != nil {
		apiPath := string(types.Endpoint_mutatingPostAuthentication)
		e.Logger.Debugf(`Registered authHook [%s]`, apiPath)
		e.POST(apiPath, func(c echo.Context) error {
			brc := c.(*types.AuthenticationHookRequest)
			out, err := authHooks.MutatingPostAuthentication(brc)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			return c.JSON(http.StatusOK, types.MiddlewareHookResponse{
				Hook:                    types.MiddlewareHook_mutatingPostAuthentication,
				Response:                out,
				SetClientRequestHeaders: HeadersToObject(brc.Request().Header),
			})
		})
	}

	if authHooks.RevalidateAuthentication != nil {
		apiPath := string(types.Endpoint_revalidateAuthentication)
		e.Logger.Debugf(`Registered authHook [%s]`, apiPath)
		e.POST(apiPath, func(c echo.Context) error {
			brc := c.(*types.AuthenticationHookRequest)
			out, err := authHooks.RevalidateAuthentication(brc)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			return c.JSON(http.StatusOK, types.MiddlewareHookResponse{
				Hook:                    types.MiddlewareHook_revalidateAuthentication,
				Response:                out,
				SetClientRequestHeaders: HeadersToObject(brc.Request().Header),
			})
		})
	}

	if authHooks.PostLogout != nil {
		apiPath := string(types.Endpoint_postLogout)
		e.Logger.Debugf(`Registered authHook [%s]`, apiPath)
		e.POST(apiPath, func(c echo.Context) error {
			brc := c.(*types.AuthenticationHookRequest)
			err := authHooks.PostLogout(brc)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			return c.JSON(http.StatusOK, types.MiddlewareHookResponse{
				Hook:                    types.MiddlewareHook_postLogout,
				SetClientRequestHeaders: HeadersToObject(brc.Request().Header),
			})
		})
	}
}
