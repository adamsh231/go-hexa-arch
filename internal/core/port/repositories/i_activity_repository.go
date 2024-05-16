package repositories

import "svc-activity/internal/core/domain/models"

type IActivityRepository interface {
	InsertActivity(model models.ActivityModel) error
	SearchActivities(service, created string, page, limit int64) (output []models.ActivityModel, err error)
	FindActivity(id string) (output models.ActivityModel, err error)
}
