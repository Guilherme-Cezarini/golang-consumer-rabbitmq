package repository

import (
	"consumer-golang/internal/domain"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
}

func NewMongoRepository(uri, dbName, messageCollection string) (*MongoRepository, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)

	collection := db.Collection(messageCollection)

	return &MongoRepository{
		client:     client,
		db:         db,
		collection: collection,
	}, nil
}

func (r *MongoRepository) SaveMessage(message *domain.Message) error {

	_, err := r.collection.InsertOne(context.TODO(), message)
	return err
}
