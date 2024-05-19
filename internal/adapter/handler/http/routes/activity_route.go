package routes

import (
	"go-hexa/internal/adapter/handler/http/handlers"

	"github.com/labstack/echo/v4"
)

func registerActivityRoute(e *echo.Group, handler handlers.Handler) {

	// group
	activityGroup := e.Group("/log")

	// routes
	activityGroup.GET("", handler.GetListActivities)
	activityGroup.GET("/:id", handler.GetDetailActivity)

}
