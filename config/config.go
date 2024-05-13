package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
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
	logFormat := os.Getenv("APP_LOG_FORMAT")
	setupLogger(logFormat)

	// Kafka
	config.Kafka.KafkaBootstrapServers = os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
	config.Kafka.KafkaTopicActivity = os.Getenv("KAFKA_TOPIC_ACTIVITY")

	// Mongo
	config.Mongo.MongoHost = os.Getenv("MONGO_HOST")
	config.Mongo.MongoPort = os.Getenv("MONGO_PORT")
	config.Mongo.MongoUser = os.Getenv("MONGO_USER")
	config.Mongo.MongoPassword = os.Getenv("MONGO_PASSWORD")

	return config, err
}

func setupLogger(format string) {

	// set log format
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "02/01/2006 15:04:05",
	})

	// set log method caller
	logrus.SetReportCaller(true)

	// set log output
	logrus.SetOutput(os.Stdout)
}
