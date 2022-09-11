package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Database {
	credential := options.Credential{
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
	}
	uri := os.Getenv("URI_DATABASE")
	clientOpts := options.Client().ApplyURI(uri).SetAuth(credential)

	ctx := context.Background()
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		panic(err)
	}
	log.Println("Successfuly connected to the database.")
	return client.Database(os.Getenv("DB_NAME"))
}
