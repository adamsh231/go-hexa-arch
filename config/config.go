package config

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"svc-activity/internal/adapter/libraries/database"
	"svc-activity/utils"
)

type Config struct {
	App   AppConfig
	Kafka KafkaConfig
	Mongo MongoConfig
}

func SetupConfig() (config Config, err error) {

	// load lib config
	if err = godotenv.Load(); err != nil {
		return config, err
	}

	// log
	setupLogger()

	// App
	config.App.Port = os.Getenv("APP_PORT")

	// Kafka
	config.Kafka.BootstrapServers = os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
	config.Kafka.GroupID = os.Getenv("KAFKA_GROUP_ID")
	config.Kafka.Topic.Activity = os.Getenv("KAFKA_TOPIC_ACTIVITY")
	if config.Kafka.WorkerPool, err = strconv.Atoi(os.Getenv("KAFKA_CONSUMER_WORKER_POOL")); err != nil {
		config.Kafka.WorkerPool = 1 // default worker pool (no pool)
	}

	// Mongo
	logrus.Info("Connecting to MongoDB!")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	user := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")
	ssl := os.Getenv("MONGO_SSL")
	srv := os.Getenv("MONGO_SRV") == "true"
	mongoLibrary := database.NewMongoLibrary(host, port, user, password, ssl, srv)
	config.Mongo.Database = os.Getenv("MONGO_DATABASE")
	config.Mongo.Client, err = mongoLibrary.Connect()
	if err != nil {
		return config, err
	}
	logrus.Info("Connected to MongoDB!")

	return config, err
}

func (c Config) CloseConfig() {

	// close mongo
	logrus.Info("disconnecting to mongo")
	if err := c.Mongo.Client.Disconnect(context.Background()); err != nil {
		logrus.Error(utils.PrintMessageWithError("error closing mongo", err))
	}
	logrus.Info("disconnected to mongo")

}

func setupLogger() {

	// set log format
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "02/01/2006 15:04:05",
	})

	// set log method caller
	logrus.SetReportCaller(true)

	// set log output
	logrus.SetOutput(os.Stdout)
}
