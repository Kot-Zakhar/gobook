package dbconnector

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var client *mongo.Client

var clientError error

var singleConnection sync.Once
var singleDisconnection sync.Once

func ConnectToDb(connectionString string) error {
	singleConnection.Do(func() {
		options := options.Client().ApplyURI(connectionString)
		client, clientError = mongo.Connect(context.TODO(), options)
		if clientError != nil {
			return
		}

		clientError = client.Ping(context.TODO(), nil)
	})

	return clientError
}

func DisconnectFromDb() {
	singleDisconnection.Do(func() {
		client.Disconnect(context.TODO())
	})
}

func GetMongoClient() (*mongo.Client, error) {
	return client, clientError
}
