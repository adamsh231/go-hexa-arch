package services

import (
	"svc-activity/internal/core/domain/entities"
)

type IActivityService interface {
	SearchActivities(input entities.SearchActivityInput) (output []entities.SearchActivityOutput, err error)
	FindActivityByID(id string) (output entities.FindActivityOutput, err error)
	InsertActivity(input entities.InsertActivityInput) (err error)
}
