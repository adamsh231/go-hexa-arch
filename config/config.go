package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Kafka KafkaConfig
	Mongo MongoConfig
}

// LoadConfig load configuration
func LoadConfig() (config Config, err error) {

	// load lib config
	if err = godotenv.Load(); err != nil {
		return config, err
	}

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
