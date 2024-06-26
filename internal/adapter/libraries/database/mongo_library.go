package database

import (
	"context"
	"fmt"
	"go-hexa/internal/core/port/libraries"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoLibrary struct {
	host     string
	port     string
	user     string
	password string
	srv      bool
	ssl      string
}

func NewMongoLibrary(host, port, user, password, ssl string, srv bool) libraries.IMongoDatabaseLibrary {
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
		uri = fmt.Sprintf("mongodb+srv://%s:%s@%s/?ssl=%s", lib.user, lib.password, lib.host, lib.ssl)
	} else {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%s/?ssl=%s", lib.user, lib.password, lib.host, lib.port, lib.ssl)
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
