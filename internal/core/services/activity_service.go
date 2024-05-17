package services

import (
	"encoding/json"
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

func (service activityService) InsertActivity(input entities.InsertActivityInput) (err error) {

	// logging
	payload, _ := json.Marshal(input)
	log := utils.NewLogUtil().LogrusWithPayload(string(payload))
	log.Info("processing insert activity")

	// construct
	now := time.Now()
	activityModel := models.ActivityModel{
		Service:   input.Service,
		Version:   input.Version,
		Message:   input.Message,
		Activity:  input.Activity,
		CreatedBy: input.CreatedBy,
		Data:      input.Data,
		Response:  input.Response,
		Created:   now,
	}
	if err = service.repo.InsertActivity(activityModel); err != nil {
		log.Error(utils.PrintMessageWithError("error insert activity", err))
		return err
	}

	// logging
	log.Info("processing insert activity, done!")

	return err
}

func (service activityService) SearchActivities(input entities.SearchActivityInput) (total int, output []entities.SearchActivityOutput, err error) {

	// logging
	payload, _ := json.Marshal(input)
	log := utils.NewLogUtil().LogrusWithPayload(string(payload))
	log.Info("processing search activity")

	// repo
	totalActivities, activities, err := service.repo.SearchActivities(input.Service, input.Created, int64(input.Page), int64(input.Limit))
	if err != nil {
		log.Error(utils.PrintMessageWithError("error while search into db", err))
		return total, output, err
	}

	// construct
	output = make([]entities.SearchActivityOutput, 0) // make it default 0 instead of null
	for _, activity := range activities {
		output = append(output, entities.SearchActivityOutput{
			ID:        activity.ID,
			Service:   activity.Service,
			Version:   activity.Version,
			Message:   activity.Message,
			Activity:  activity.Activity,
			CreatedBy: activity.CreatedBy,
			Created:   activity.Created.Format(time.DateTime),
		})
	}

	// logging
	log.Info("processing search activity, done!")

	return int(totalActivities), output, err
}

func (service activityService) FindActivityByID(id string) (output entities.FindActivityOutput, err error) {

	// logging
	log := utils.NewLogUtil().LogrusWithPayload(id)
	log.Info("processing search activity")

	// repo
	activity, err := service.repo.FindActivity(id)
	if err != nil {
		log.Error(utils.PrintMessageWithError("error while search into db", err))
		return output, err
	}

	// construct
	output = entities.FindActivityOutput{
		ID:        activity.ID,
		Service:   activity.Service,
		Version:   activity.Version,
		Message:   activity.Message,
		Activity:  activity.Activity,
		CreatedBy: activity.CreatedBy,
		Data:      activity.Data,
		Response:  activity.Response,
		Created:   activity.Created.Format(time.DateTime),
	}

	// logging
	log.Info("processing search activity, done!")

	return output, err
}
