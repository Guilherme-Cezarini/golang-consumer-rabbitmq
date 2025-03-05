package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	MongoDBDatabase        string
	MongoDBUsername        string
	MongoDBPassword        string
	MongoMessageCollection string
	MongoDBUrl             string
	RabbitMQUrl            string
	RabbitMQUsername       string
	RabbitMQPassword       string
	RabbitMQQueue          string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	return &Config{
		RabbitMQUsername:       os.Getenv("RABBIT_MQ_USERNAME"),
		RabbitMQPassword:       os.Getenv("RABBIT_MQ_PASSWORD"),
		RabbitMQQueue:          os.Getenv("RABBIT_MQ_QUEUE"),
		RabbitMQUrl:            os.Getenv("RABBIT_MQ_URL"),
		MongoMessageCollection: os.Getenv("MONGO_MESSAGE_COLLECTION"),
		MongoDBUsername:        os.Getenv("MONGO_USERNAME"),
		MongoDBDatabase:        os.Getenv("MONGO_DATABASE"),
		MongoDBPassword:        os.Getenv("MONGO_PASSWORD"),
		MongoDBUrl:             os.Getenv("MONGO_URL"),
	}, nil
}
