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

func DBSetTest() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	MongoDb := os.Getenv("MONGODB_URL_TEST")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MongoDb))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("failed to connect to test mongodb")
		return nil
	}

	fmt.Println("Successfully connected to test mongodb")

	return client
}

var ClientTest *mongo.Client = DBSetTest()

func UserTestData(user *mongo.Client, collectionName string) *mongo.Collection {
	var policyCollection *mongo.Collection = user.Database("User").Collection(collectionName)
	return policyCollection
}

func RestaurantTestData(restaurants *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = restaurants.Database("Restaurants").Collection(collectionName)
	return collection
}

func OrderTestData(employee *mongo.Client, collectionName string) *mongo.Collection {
	var employeeCollection *mongo.Collection = employee.Database("Orders").Collection(collectionName)
	return employeeCollection
}
