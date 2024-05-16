package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"svc-activity/internal/core/domain/models"
	"svc-activity/internal/core/port/repositories"
)

type activityRepository struct {
	client   *mongo.Client
	database string
}

func NewActivityRepository(client *mongo.Client, database string) repositories.IActivityRepository {
	return activityRepository{client: client, database: database}
}

func (repo activityRepository) getCollection() *mongo.Collection{
	return repo.client.Database(repo.database).Collection(models.ActivityCollection)
}

func (repo activityRepository) InsertActivity(model models.InsertActivityModel) (err error) {
	if _, err = repo.getCollection().InsertOne(context.Background(), model); err != nil{
		return err
	}
	return err
}
