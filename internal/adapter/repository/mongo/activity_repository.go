package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"svc-activity/internal/core/domain/models"
	"svc-activity/internal/core/port/repositories"
	"time"
)

type activityRepository struct {
	client   *mongo.Client
	database string
}

func NewActivityRepository(client *mongo.Client, database string) repositories.IActivityRepository {
	return activityRepository{client: client, database: database}
}

func (repo activityRepository) getCollection() *mongo.Collection {
	return repo.client.Database(repo.database).Collection(models.ActivityCollection)
}

func (repo activityRepository) InsertActivity(model models.ActivityModel) (err error) {
	if _, err = repo.getCollection().InsertOne(context.Background(), model); err != nil {
		return err
	}
	return err
}

func (repo activityRepository) SearchActivities(service, created string, page, limit int64) (output []models.ActivityModel, err error) {

	// filter
	filter := bson.M{
		"created": bson.M{
			"$gte": time.Now().Add(-1 * 24 * time.Hour),
			"$lte": time.Now().Add(24 * time.Hour),
		},
	}

	// filter service
	if service != ""{
		filter["service"] = service
	}

	// find
	offset := (page - 1) * limit
	cursor, err := repo.getCollection().Find(context.Background(), filter, &options.FindOptions{
		Limit: &limit,
		Skip:  &offset,
	})
	if err != nil {
		return output, err
	}
	defer cursor.Close(context.Background())

	// cursors
	if err := cursor.All(context.Background(), &output); err != nil {
		return output, err
	}

	return output, err
}

func (repo activityRepository) FindActivity(id string) (output models.ActivityModel, err error) {

	// filter
	filter := bson.M{
		"id": id,
	}

	// Construct your query
	if err = repo.getCollection().FindOne(context.Background(), filter).Decode(&output); err != nil {
		return output, err
	}

	return output, err
}
