package config

import (
	"go-hexa/internal/adapter/repository/mongo"
	services2 "go-hexa/internal/core/port/services"
	"go-hexa/internal/core/services"
)

type ServiceInjector struct {
	ActivityService services2.IActivityService
}

func InitInjection(config Config) (inject ServiceInjector) {

	// activity
	activityRepo := mongo.NewActivityRepository(config.Mongo.Client, config.Mongo.Database)
	inject.ActivityService = services.NewActivityService(activityRepo)

	return inject
}
