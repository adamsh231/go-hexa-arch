package handlers

import (
	"go-hexa/config"
	"go-hexa/internal/core/domain/entities"
	"go-hexa/internal/core/domain/presenters"
	"go-hexa/utils"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	injector config.ServiceInjector
}

func NewHandler(injector config.ServiceInjector) Handler {
	return Handler{injector: injector}
}

// GetListActivities
//
//	@Summary					Search log activities
//	@Tags						Logging
//	@Accept						json
//	@Produce					json
//	@SecurityDefinitions.apikey	ApiKeyAuth
//	@Param						page	query		presenters.GetListActivitiesRequest	true	"query"
//	@Success					200		{object}	utils.ResponseUtil{data=[]entities.SearchActivityOutput,meta=utils.ResponseMetaUtil}
//	@Failure					400		{object}	utils.ResponseStatusUtil
//	@Failure					422		{object}	utils.ResponseStatusUtil
//	@Failure					500		{object}	utils.ResponseStatusUtil
//	@Router						/v1/log [GET]
func (handler Handler) GetListActivities(c echo.Context) error {

	// bind queries
	var request presenters.GetListActivitiesRequest
	err := c.Bind(&request)
	if err != nil {
		return utils.ResponseError(c, http.StatusBadRequest, err.Error())
	}

	// validation
	if errs := utils.ValidateStruct(request); len(errs) > 0 {
		return utils.ResponseListError(c, http.StatusBadRequest, "validation Error", errs)
	}
	parseDate, err := time.Parse(time.DateOnly, request.Date)
	if err != nil {
		return utils.ResponseError(c, http.StatusUnprocessableEntity, err.Error())
	}

	// construct
	page, limit := utils.ValidationPaginationDefault(request.Page, request.Limit)
	input := entities.SearchActivityInput{
		Service: request.Service,
		Created: parseDate,
		Page:    page,
		Limit:   limit,
	}

	// service
	totalActivities, activities, err := handler.injector.ActivityService.SearchActivities(input)
	if err != nil {
		return utils.ResponseError(c, http.StatusInternalServerError, err.Error())
	}
	meta := utils.CreateMetaPagination(page, limit, totalActivities)

	return utils.ResponseSuccessDataWithMeta(c, http.StatusOK, activities, meta, "")
}

// GetDetailActivity
//
//	@Summary					Search log activities
//	@Tags						Logging
//	@Accept						json
//	@Produce					json
//	@SecurityDefinitions.apikey	ApiKeyAuth
//	@Param						id	path		string	true	"id"
//	@Success					200	{object}	utils.ResponseUtil{data=entities.FindActivityOutput}
//	@Failure					404	{object}	utils.ResponseStatusUtil
//	@Failure					500	{object}	utils.ResponseStatusUtil
//	@Router						/v1/log/{id} [GET]
func (handler Handler) GetDetailActivity(c echo.Context) error {

	// param
	id := c.Param("id")

	// service
	activity, err := handler.injector.ActivityService.FindActivityByID(id)
	if err != nil {
		return utils.ResponseError(c, http.StatusInternalServerError, err.Error())
	}

	// validate
	if activity.ID == "" {
		return utils.ResponseError(c, http.StatusNotFound, "data not found")
	}

	return utils.ResponseSuccessData(c, http.StatusOK, activity, "")
}
