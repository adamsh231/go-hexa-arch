package libraries

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoLibrary struct {
	Host     string
	Port     string
	User     string
	Password string
	SRV      bool
	SSL      string
}

func (lib MongoLibrary) Connect() (client *mongo.Client, err error) {

	// Set client options
	var uri string
	if lib.SRV {
		uri = fmt.Sprintf("mongodb+srv://%s/?ssl=%s", lib.Host, lib.SSL)
	} else {
		uri = fmt.Sprintf("mongodb://%s:%s/?ssl=%s", lib.Host, lib.Port, lib.SSL)
	}
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return client, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return client, err
	}

	return client, err
}
