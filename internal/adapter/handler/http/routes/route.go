package routes

import (
	"github.com/labstack/echo/v4"
	"svc-activity/internal/adapter/handler/http/handlers"
)

// RegisterRoute All sub group route must be registered here
func RegisterRoute(e *echo.Echo, handler handlers.Handler) {
	registerActivityRoute(e, handler)
}
