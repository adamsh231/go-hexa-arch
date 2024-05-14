package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"svc-activity/internal/core/port/libraries"
)

type mongoLibrary struct {
	host     string
	port     string
	user     string
	password string
	srv      bool
	ssl      string
}

func NewMongoLibrary(host, port, user, password, ssl string, srv bool) libraries.IMongoDatabaseLibrary{
	return mongoLibrary{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		srv:      srv,
		ssl:      ssl,
	}
}

func (lib mongoLibrary) Connect() (client *mongo.Client, err error) {

	// Set client options
	var uri string
	if lib.srv {
		uri = fmt.Sprintf("mongodb+srv://%s/?ssl=%s", lib.host, lib.ssl)
	} else {
		uri = fmt.Sprintf("mongodb://%s:%s/?ssl=%s", lib.host, lib.port, lib.ssl)
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