package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"svc-activity/libraries"
)

type Config struct {
	Kafka KafkaConfig
	Mongo MongoConfig
}

// SetupConfig load configuration
func SetupConfig() (config Config, err error) {

	// load lib config
	if err = godotenv.Load(); err != nil {
		return config, err
	}

	// log
	setupLogger()

	// Kafka
	config.Kafka.BootstrapServers = os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
	config.Kafka.GroupID = os.Getenv("KAFKA_GROUP_ID")
	config.Kafka.Topic.Activity = os.Getenv("KAFKA_TOPIC_ACTIVITY")
	if config.Kafka.WorkerPool, err = strconv.Atoi(os.Getenv("KAFKA_CONSUMER_WORKER_POOL")); err != nil {
		config.Kafka.WorkerPool = 1 // default worker pool (no pool)
	}

	// Mongo
	logrus.Info("Connecting to MongoDB!")
	srv := os.Getenv("MONGO_SRV") == "true"
	mongoLibrary := libraries.MongoLibrary{
		Host:     os.Getenv("MONGO_HOST"),
		Port:     os.Getenv("MONGO_PORT"),
		User:     os.Getenv("MONGO_USER"),
		Password: os.Getenv("MONGO_PASSWORD"),
		SSL:      os.Getenv("MONGO_SSL"),
		SRV:      srv,
	}
	config.Mongo.Database = os.Getenv("MONGO_DATABASE")
	config.Mongo.Client, err = mongoLibrary.Connect()
	if err != nil {
		return config, err
	}
	logrus.Info("Connected to MongoDB!")

	return config, err
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
