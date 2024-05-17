package routes

import (
	"github.com/labstack/echo/v4"
	"svc-activity/config"
	"svc-activity/internal/adapter/handler/http/handlers"
	"svc-activity/internal/adapter/handler/http/middlewares"
)

// RegisterRoute All subgroup route must be registered here
func RegisterRoute(e *echo.Echo, handler handlers.Handler, apiKeyConfig config.ApiKeyConfig) {

	// v1 handlers
	v1 := e.Group("/v1", middlewares.ApiKeyMiddleware(apiKeyConfig.General))
	registerActivityRoute(v1, handler)
}
