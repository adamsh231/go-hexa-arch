package services

import (
	"fmt"
	"svc-activity/internal/core/domain/entities"
	"svc-activity/internal/core/domain/models"
	"svc-activity/internal/core/port/services/activity"
)

type activityService struct {
	repo activity.IActivityRepository
}

func NewActivityService(repo activity.IActivityRepository) activity.IActivityService{
	return activityService{repo: repo}
}

func (service activityService) InsertActivity(input entities.InsertActivityInput) {
	if err := service.repo.InsertActivity(models.InsertActivityModel{}); err != nil {
		return
	}
	fmt.Println("ini service")
}