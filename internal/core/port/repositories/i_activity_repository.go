package repositories

import (
	"svc-activity/internal/core/domain/models"
	"time"
)

type IActivityRepository interface {
	InsertActivity(model models.ActivityModel) error
	SearchActivities(service string, created time.Time, page, limit int64) (total int64, output []models.ActivityModel, err error)
	FindActivity(id string) (output models.ActivityModel, err error)
}
