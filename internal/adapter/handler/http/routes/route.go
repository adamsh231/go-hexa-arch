package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-hexa/config"
	"go-hexa/internal/adapter/handler/http/handlers"
)

// RegisterRoute All subgroup route must be registered here
func RegisterRoute(e *fiber.App, handler handlers.Handler, apiKeyConfig config.ApiKeyConfig) {

	// v1 handlers
	v1 := e.Group("/v1")
	registerActivityRoute(v1, handler)
}
