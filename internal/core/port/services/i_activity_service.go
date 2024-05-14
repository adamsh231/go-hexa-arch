package services

import (
	"svc-activity/internal/core/domain/entities"
)

type IActivityService interface {
	InsertActivity(input entities.InsertActivityInput)
}