package services

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"svc-activity/internal/core/domain/entities"
	"svc-activity/internal/core/domain/models"
	"svc-activity/internal/core/port/repositories"
	"svc-activity/internal/core/port/services"
	"svc-activity/utils"
	"time"
)

type activityService struct {
	repo repositories.IActivityRepository
}

func NewActivityService(repo repositories.IActivityRepository) services.IActivityService {
	return activityService{repo: repo}
}

func (service activityService) InsertActivity(input entities.InsertActivityInput) {

	// logging
	payload, _ := json.Marshal(input)
	utils.LogrusWithPayload(string(payload)).Info("processing insert activity")

	// construct
	now := time.Now()
	activityModel := models.InsertActivityModel{
		Service:   input.Service,
		Version:   input.Version,
		Message:   input.Message,
		Activity:  input.Activity,
		CreatedBy: input.CreatedBy,
		Data:      input.Data,
		Response:  input.Response,
		Created:   now,
	}
	if err := service.repo.InsertActivity(activityModel); err != nil {
		utils.LogrusWithPayload(string(payload)).Error(err)
		return
	}

	// logging
	logrus.WithField("payload", string(payload)).Info("processing insert activity, done!")
}