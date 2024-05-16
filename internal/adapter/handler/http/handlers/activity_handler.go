package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"svc-activity/config"
)

type Handler struct {
	injector config.ServiceInjector
}

func NewHandler(injector config.ServiceInjector) Handler {
	return Handler{injector: injector}
}

func (handler Handler) GetListActivities(c echo.Context) error {
	return c.String(http.StatusOK, "GetListActivities")
}

func (handler Handler) GetDetailActivity(c echo.Context) error {
	id := c.Param("id")
	fmt.Println(id)
	return c.String(http.StatusOK, fmt.Sprintf("GetDetailActivity %s", id))
}