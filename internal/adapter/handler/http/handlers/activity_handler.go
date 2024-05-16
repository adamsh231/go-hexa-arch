package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"svc-activity/config"
	"svc-activity/internal/core/domain/entities"
	"svc-activity/internal/core/domain/presenters"
	"svc-activity/utils"
)

type Handler struct {
	injector config.ServiceInjector
}

func NewHandler(injector config.ServiceInjector) Handler {
	return Handler{injector: injector}
}

func (handler Handler) GetListActivities(c echo.Context) error {

	// bind queries
	var request presenters.GetListActivitiesRequest
	err := c.Bind(&request); if err != nil {
		return utils.ResponseError(c, http.StatusBadRequest, err.Error())
	}

	// construct
	input := entities.SearchActivityInput{
		Service: request.Service,
		Created: request.Created,
		Page:    request.Page,
		Limit:   request.Limit,
	}

	// service
	activity, err := handler.injector.ActivityService.SearchActivities(input)
	if err != nil{
		return utils.ResponseError(c, http.StatusInternalServerError, err.Error())
	}

	return utils.ResponseSuccessDataWithMeta(c, http.StatusOK, activity, "", "")
}

func (handler Handler) GetDetailActivity(c echo.Context) error {

	// param
	id := c.Param("id")

	// service
	activity, err := handler.injector.ActivityService.FindActivityByID(id)
	if err != nil{
		return utils.ResponseError(c, http.StatusInternalServerError, err.Error())
	}

	return utils.ResponseSuccessData(c, http.StatusOK, activity, "")
}