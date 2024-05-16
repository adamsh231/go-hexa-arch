package routes

import (
	"github.com/labstack/echo/v4"
	"svc-activity/internal/adapter/handler/http/handlers"
)

func registerActivityRoute(e *echo.Echo, handler handlers.Handler){

	// group
	activityGroup := e.Group("/log")

	// routes
	activityGroup.GET("", handler.GetListActivities)
	activityGroup.GET("/:id", handler.GetDetailActivity)

}