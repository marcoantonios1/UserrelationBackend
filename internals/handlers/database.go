package handlers

import (
	"os"
	"userrelation/pkg/database"

	"go.mongodb.org/mongo-driver/mongo"
)

var Neo4j_User = os.Getenv("NEO4J_USER")

func Neo4j(env string) string {
	var client string

	if env == "prod" {
		client = os.Getenv("NEO4J_URL_USER")
	} else {
		client = os.Getenv("NEO4J_URL_USER_TEST")
	}

	return client
}

func Neo4j_Password(env string) string {
	var client string

	if env == "prod" {
		client = os.Getenv("NEO4J_PASSWORD")
	} else {
		client = os.Getenv("NEO4J_PASSWORD_TEST")
	}

	return client
}

func Neo4j_Database(env string) string {
	var client string

	if env == "prod" {
		client = os.Getenv("NEO4J_DATABASE")
	} else {
		client = os.Getenv("NEO4J_DATABASE_TEST")
	}

	return client
}

func FeedbackCollection(env string) *mongo.Collection {
	var client *mongo.Collection

	if env == "prod" {
		client = database.OrdersData(database.Orders, "Feedbacks")
	} else {
		client = database.OrderTestData(database.ClientTest, "Feedbacks")
	}

	return client
}

func LocationCollection(env string) *mongo.Collection {
	var client *mongo.Collection

	if env == "prod" {
		client = database.RestaurantsData(database.Restaurants, "Location")
	} else {
		client = database.RestaurantTestData(database.ClientTest, "Location")
	}

	return client
}

func RestaurantCollection(env string) *mongo.Collection {
	var client *mongo.Collection

	if env == "prod" {
		client = database.RestaurantsData(database.Restaurants, "Restaurants")
	} else {
		client = database.RestaurantTestData(database.ClientTest, "Restaurants")
	}

	return client
}

func UsersCollection(env string) *mongo.Collection {
	var client *mongo.Collection

	if env == "prod" {
		client = database.UsersData(database.Users, "Users")
	} else {
		client = database.UserTestData(database.ClientTest, "Users")
	}

	return client
}
