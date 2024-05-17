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

func CreateMetaPagination() map[string]interface{} {
	return map[string]interface{}{
		"current_page": 1,
		"per_page":     2,
		"from":         3,
		"to":           4,
		"total":        5,
		"last_page":    6,
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
