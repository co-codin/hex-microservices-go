package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)


type mongoRepository struct {
	client *mongo.Client
	database string
	timeout time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	return nil, nil
}