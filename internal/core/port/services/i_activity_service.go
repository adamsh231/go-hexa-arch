package services

import (
	"go-hexa/internal/core/domain/entities"
)

type IActivityService interface {
	SearchActivities(input entities.SearchActivityInput) (total int, output []entities.SearchActivityOutput, err error)
	FindActivityByID(id string) (output entities.FindActivityOutput, err error)
	InsertActivity(input entities.InsertActivityInput) (err error)
}
