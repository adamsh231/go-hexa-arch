package config

import (
	"svc-activity/internal/adapter/repository/mongo"
	services2 "svc-activity/internal/core/port/services"
	"svc-activity/internal/core/services"
)

type ServiceInjector struct {
	ActivityService services2.IActivityService
}

func InitInjection(config Config) (inject ServiceInjector){

	// activity
	activityRepo := mongo.NewActivityRepository(config.Mongo.Client)
	inject.ActivityService = services.NewActivityService(activityRepo)


	return inject
}
