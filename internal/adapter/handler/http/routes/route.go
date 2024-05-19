package routes

import (
	"go-hexa/config"
	"go-hexa/internal/adapter/handler/http/handlers"
	"go-hexa/internal/adapter/handler/http/middlewares"

	"github.com/labstack/echo/v4"
)

// RegisterRoute All subgroup route must be registered here
func RegisterRoute(e *echo.Echo, handler handlers.Handler, apiKeyConfig config.ApiKeyConfig) {

	// v1 handlers
	v1 := e.Group("/v1", middlewares.ApiKeyMiddleware(apiKeyConfig.General))
	registerActivityRoute(v1, handler)
}
