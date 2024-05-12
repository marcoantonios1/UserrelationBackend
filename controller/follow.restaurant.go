package controller

import (
	"context"
	"net/http"
	"userrelation/database"
	"userrelation/helper"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var RestaurantCollection *mongo.Collection = database.RestaurantsData(database.Restaurants, "Restaurants")

func UnFollowRestaurant() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from context
		userID, exists := c.Get("id")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
			return
		}

		userIDObj := userID.(primitive.ObjectID)

		userToUnFollowID := c.Query("resto_id")
		userToUnFollowIDObj, err := primitive.ObjectIDFromHex(userToUnFollowID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID to unfollow"})
			return
		}

		ctx := context.Background() // Consider using request-scoped context

		// Decrement the 'following' count of the user
		update := bson.M{"$inc": bson.M{"following_restaurants": -1}}
		_, err = UsersCollection.UpdateOne(ctx, bson.M{"_id": userIDObj}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		// Decrement the 'followers' count of the user being unfollowed
		update = bson.M{"$inc": bson.M{"followers": -1}}
		_, err = RestaurantCollection.UpdateOne(ctx, bson.M{"_id": userToUnFollowIDObj}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user to unfollow"})
			return
		}

		c.JSON(http.StatusOK, "FOLLOW")
		go helper.KafkaUnFollowRestaurant(ctx, userIDObj.Hex(), userToUnFollowID)
	}
}

func FollowRestaurant() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from context
		userID, exists := c.Get("id")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
			return
		}

		userIDObj := userID.(primitive.ObjectID)

		userToFollowID := c.Query("resto_id")
		userToFollowIDObj, err := primitive.ObjectIDFromHex(userToFollowID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID to follow"})
			return
		}

		ctx := context.Background() // Consider using request-scoped context

		// Increment the 'following' count of the user
		update := bson.M{"$inc": bson.M{"following_restaurants": 1}}
		_, err = UsersCollection.UpdateOne(ctx, bson.M{"_id": userIDObj}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		// Increment the 'followers' count of the user being followed
		update = bson.M{"$inc": bson.M{"followers": 1}}
		_, err = RestaurantCollection.UpdateOne(ctx, bson.M{"_id": userToFollowIDObj}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user to follow"})
			return
		}

		c.JSON(http.StatusOK, "FOLLOWING")
		go helper.KafkaFollowRestaurant(ctx, userIDObj.Hex(), userToFollowID)
	}
}
