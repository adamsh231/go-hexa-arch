package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"svc-activity/utils"
)

//func ApiKeyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		authKey := c.Request().Header.Get("x-api-key")
//		if authKey == "" {
//			return utils.ResponseError(c, http.StatusUnauthorized, "Unauthorized")
//		}
//		return next(c)
//	}
//}

func ApiKeyMiddleware(apiKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authKey := c.Request().Header.Get("x-api-key")
			if authKey == "" {
				return utils.ResponseError(c, http.StatusUnauthorized, "x-api-key required on request header")
			}
			if authKey != apiKey {
				return utils.ResponseError(c, http.StatusUnauthorized, "Unauthorized")
			}
			return next(c)
		}
	}
}
