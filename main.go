package main

import (
	"consumer-golang/internal/config"
	"consumer-golang/internal/delivery"
	"consumer-golang/internal/repository"
	"consumer-golang/internal/service"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	mongoUri := fmt.Sprintf("mongodb://%s:%s@%s", cfg.MongoDBUsername, cfg.MongoDBPassword, cfg.MongoDBUrl)

	mongoRepo, err := repository.NewMongoRepository(mongoUri, cfg.MongoDBDatabase, cfg.MongoMessageCollection)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v", err)
	}

	messageService := service.NewMessageService(mongoRepo)

	rabbitMQUri := fmt.Sprintf("amqp://%s:%s@%s", cfg.RabbitMQUsername, cfg.RabbitMQPassword, cfg.RabbitMQUrl)
	rabbitMQConsumer, err := delivery.NewRabbitMQConsumer(rabbitMQUri, messageService, cfg.RabbitMQQueue)
	if err != nil {
		log.Fatalf("Could not connect to RabbitMQ: %v", err)
	}

	log.Println("Starting RabbitMQ Consumer...")
	rabbitMQConsumer.Consume()
}
