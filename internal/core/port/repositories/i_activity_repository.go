package repositories

import "svc-activity/internal/core/domain/models"

type IActivityRepository interface {
	InsertActivity(model models.InsertActivityModel) error
}
