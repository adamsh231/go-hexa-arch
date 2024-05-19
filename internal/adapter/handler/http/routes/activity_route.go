package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-hexa/internal/adapter/handler/http/handlers"
)

func registerActivityRoute(e fiber.Router, handler handlers.Handler) {

	// group
	activityGroup := e.Group("/log")

	// routes
	activityGroup.Get("", handler.GetListActivities)
	activityGroup.Get("/:id", handler.GetDetailActivity)

}
