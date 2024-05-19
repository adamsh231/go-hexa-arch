package mongo

import (
	"context"
	"go-hexa/internal/core/domain/models"
	"go-hexa/internal/core/port/repositories"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (repo activityRepository) SearchActivities(service string, created time.Time, page, limit int64) (total int64, output []models.ActivityModel, err error) {

	// filter
	filter := bson.M{
		"created": bson.M{
			"$gte": created,
			"$lte": created.Add(24 * time.Hour),
		},
	}

	// filter service
	if service != "" {
		filter["service"] = service
	}

	// find
	offset := (page - 1) * limit
	cursor, err := repo.getCollection().Find(context.Background(), filter, &options.FindOptions{
		Limit: &limit,
		Skip:  &offset,
	})
	if err != nil {
		return total, output, err
	}
	defer cursor.Close(context.Background())

	// get total
	total, err = repo.getCollection().CountDocuments(context.Background(), filter)
	if err != nil {
		return total, output, err
	}

	// cursors
	if err := cursor.All(context.Background(), &output); err != nil {
		return total, output, err
	}

	return total, output, err
}

func (repo activityRepository) FindActivity(id string) (output models.ActivityModel, err error) {

	// filter
	idPrimitive, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": idPrimitive,
	}

	// Construct your query
	if err = repo.getCollection().FindOne(context.Background(), filter).Decode(&output); err == mongo.ErrNoDocuments {
		return output, nil
	} else if err != nil {
		return output, err
	}

	return output, err
}
