package mongo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"svc-activity/internal/core/domain/models"
	"svc-activity/internal/core/port/services/activity"
)

type activityRepository struct {
	client *mongo.Client
}

func NewActivityRepository(client *mongo.Client) activity.IActivityRepository {
	return activityRepository{client: client}
}

func (repo activityRepository) InsertActivity(model models.InsertActivityModel) error {
	fmt.Println("ini repo")
	return nil
}
