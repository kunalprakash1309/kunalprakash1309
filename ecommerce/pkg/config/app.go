package config

import (
	"context"
	"time"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// Setup initializes a mongo client
func Setup(ctx context.Context, address string) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(address))
	if err != nil {
		panic(err.Error())
	}

	err = client.Connect(ctx)
	if err != nil {
		panic(err.Error())
	}

	db := client.Database("Ecommerce")
	log.Println("Database connected")

	DB = db
	
}
