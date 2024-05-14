package config

import (
	"svc-activity/internal/adapter/repository/mongo"
	"svc-activity/internal/core/port/services/activity"
	"svc-activity/internal/core/services"
)

type ServiceInjector struct {
	ActivityService activity.IActivityService
}

func InitInjection(config Config) (inject ServiceInjector){

	// activity
	activityRepo := mongo.NewActivityRepository(config.Mongo.Client)
	inject.ActivityService = services.NewActivityService(activityRepo)


	return inject
}
