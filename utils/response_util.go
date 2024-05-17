package utils

import (
	"github.com/labstack/echo/v4"
)

type ResponseUtil struct {
	Status *ResponseStatusUtil `json:"status,omitempty"`
	Data   interface{}         `json:"data,omitempty"`
	Meta   interface{}         `json:"meta,omitempty"`
}

type ResponseStatusUtil struct {
	Code    string   `json:"code,omitempty"`
	Message string   `json:"message,omitempty"`
	Errors  []string `json:"errors,omitempty"`
}

type ResponseMetaUtil struct {
	CurrentPage int `json:"current_page,omitempty"`
	PerPage     int `json:"per_page,omitempty"`
	From        int `json:"from,omitempty"`
	To          int `json:"to,omitempty"`
	Total       int `json:"total,omitempty"`
	LastPage    int `json:"last_page,omitempty"`
}

func CreateMetaPagination(page, limit, total int) ResponseMetaUtil {

	// last page
	lastPage := (total + limit - 1) / limit
	if page > lastPage {
		page = lastPage
	}

	// from and to
	from := (page-1)*limit + 1
	to := from + limit - 1
	if to > total {
		to = total
	}

	return ResponseMetaUtil{
		CurrentPage: page,
		PerPage:     limit,
		From:        from,
		To:          to,
		Total:       total,
		LastPage:    lastPage,
	}
}

func ResponseError(c echo.Context, statusCode int, message string) error {
	response := ResponseUtil{
		Status: &ResponseStatusUtil{
			Message: message,
		},
	}
	return c.JSON(statusCode, response)
}

func ResponseListError(c echo.Context, statusCode int, message string, errs []error) error {

	// fill error
	var errors []string
	for _, err := range errs {
		errors = append(errors, err.Error())
	}

	// response
	response := ResponseUtil{
		Status: &ResponseStatusUtil{
			Message: message,
			Errors:  errors,
		},
	}

	return c.JSON(statusCode, response)
}

func ResponseSuccess(c echo.Context, statusCode int, message string) error {
	response := ResponseUtil{
		Status: &ResponseStatusUtil{
			Message: message,
		},
	}
	return c.JSON(statusCode, response)
}

func ResponseSuccessData(c echo.Context, statusCode int, data interface{}, message string) error {
	response := ResponseUtil{
		Data: data,
	}
	if message != "" {
		response.Status = &ResponseStatusUtil{
			Message: message,
		}
	}
	return c.JSON(statusCode, response)
}

func ResponseSuccessDataWithMeta(c echo.Context, statusCode int, data, meta interface{}, message string) error {
	response := ResponseUtil{
		Data: data,
		Meta: meta,
	}
	if message != "" {
		response.Status = &ResponseStatusUtil{
			Message: message,
		}
	}
	return c.JSON(statusCode, response)
}
