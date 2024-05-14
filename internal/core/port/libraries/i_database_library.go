package libraries

import "go.mongodb.org/mongo-driver/mongo"

type IMongoDatabaseLibrary interface {
	Connect() (client *mongo.Client, err error)
}