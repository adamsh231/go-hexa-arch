package config

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type AppConfig struct {
	Port string
}

type KafkaConfig struct {
	BootstrapServers string
	GroupID          string
	Topic            KafkaTopic
	WorkerPool       int
}

type KafkaTopic struct {
	Activity string
}

type MongoConfig struct {
	Client   *mongo.Client
	Database string
}
