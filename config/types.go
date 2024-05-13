package config

type KafkaConfig struct {
	KafkaBootstrapServers string
	KafkaTopicActivity    string
}

type MongoConfig struct {
	MongoHost     string
	MongoPort     string
	MongoUser     string
	MongoPassword string
}