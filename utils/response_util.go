package utils

import "github.com/labstack/echo/v4"

type ResponseUtil struct {
	Status ResponseStatusUtil `json:"status,omitempty"`
	Data   interface{}        `json:"data,omitempty"`
	Meta   interface{}        `json:"meta,omitempty"`
}

type ResponseStatusUtil struct {
	Code    string      `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
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
	return c.JSON(statusCode, ResponseUtil{
		Status: ResponseStatusUtil{
			Message: message,
		},
	})
}

func ResponseSuccess(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, ResponseUtil{
		Status: ResponseStatusUtil{
			Message: message,
		},
	})
}

func ResponseSuccessData(c echo.Context, statusCode int, data interface{}, message string) error {
	return c.JSON(statusCode, ResponseUtil{
		Status: ResponseStatusUtil{
			Message: message,
		},
		Data: data,
	})
}


func ResponseSuccessDataWithMeta(c echo.Context, statusCode int, data, meta interface{}, message string) error {
	return c.JSON(statusCode, ResponseUtil{
		Status: ResponseStatusUtil{
			Message: message,
		},
		Data: data,
		Meta: meta,
	})
}


