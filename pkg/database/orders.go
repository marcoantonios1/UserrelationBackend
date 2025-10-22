package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSetOrders() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	MongoDb := os.Getenv("MONGODB_URL_ORDERS")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MongoDb))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("failed to connect to mongodb orders")
		return nil
	}

	fmt.Println("Successfully connected to mongodb orders")

	return client
}

var Orders *mongo.Client = DBSetOrders()

func OrdersData(orders *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = orders.Database("Orders").Collection(collectionName)
	return collection
}
