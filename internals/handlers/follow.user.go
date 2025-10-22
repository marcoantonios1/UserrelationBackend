package handlers

import (
	"context"
	"net/http"
	"userrelation/internals/models"
	helper "userrelation/internals/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTotalFollowRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		environement := c.GetString("env")
		// var prod bool
		// if environement == "prod" {
		// 	prod = true
		// } else {
		// 	prod = false
		// }
		// Get user ID from context
		userID, exists := c.Get("id")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
			return
		}

		userIDObj := userID.(primitive.ObjectID)

		ctx := context.Background() // Consider using request-scoped context

		// Find the user
		var user models.User
		err := UsersCollection(environement).FindOne(ctx, bson.M{"_id": userIDObj}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
			return
		}

		// Return the Follow_Request field of the user
		c.JSON(http.StatusOK, user.Follow_Request)
	}
}

func Follow() gin.HandlerFunc {
	return func(c *gin.Context) {
		environement := c.GetString("env")
		var prod bool
		if environement == "prod" {
			prod = true
		} else {
			prod = false
		}

		// Get user ID from context
		userID, exists := c.Get("id")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
			return
		}

		userIDObj, ok := userID.(primitive.ObjectID)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		userToFollowID := c.Query("user_id")
		userToFollowIDObj, err := primitive.ObjectIDFromHex(userToFollowID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID to follow"})
			return
		}

		if userIDObj == userToFollowIDObj {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot follow yourself"})
			return
		}

		ctx := context.Background() // Consider using request-scoped context

		// Increment the 'following' count of the user
		update := bson.M{"$inc": bson.M{"following": 1}}
		_, err = UsersCollection(environement).UpdateOne(ctx, bson.M{"_id": userIDObj}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		// Increment the 'followers' count of the user being followed
		update = bson.M{"$inc": bson.M{"followers": 1}}
		_, err = UsersCollection(environement).UpdateOne(ctx, bson.M{"_id": userToFollowIDObj}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user to follow"})
			return
		}

		c.JSON(http.StatusOK, "FOLLOWING")
		go helper.KafkaFollow(ctx, userIDObj.Hex(), userToFollowID, prod)
		go helper.KafkaFollowLog(ctx, userIDObj.Hex(), userToFollowID, "followed", prod)
	}
}

func UnFollow() gin.HandlerFunc {
	return func(c *gin.Context) {
		environement := c.GetString("env")
		var prod bool
		if environement == "prod" {
			prod = true
		} else {
			prod = false
		}

		// Get user ID from context
		userID, exists := c.Get("id")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
			return
		}

		userIDObj, ok := userID.(primitive.ObjectID)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		userToUnFollowID := c.Query("user_id")
		userToUnFollowIDObj, err := primitive.ObjectIDFromHex(userToUnFollowID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID to unfollow"})
			return
		}

		ctx := context.Background() // Consider using request-scoped context

		// Decrement the 'following' count of the user
		update := bson.M{"$inc": bson.M{"following": -1}}
		_, err = UsersCollection(environement).UpdateOne(ctx, bson.M{"_id": userIDObj}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		// Decrement the 'followers' count of the user being unfollowed
		update = bson.M{"$inc": bson.M{"followers": -1}}
		_, err = UsersCollection(environement).UpdateOne(ctx, bson.M{"_id": userToUnFollowIDObj}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user to unfollow"})
			return
		}

		c.JSON(http.StatusOK, "FOLLOW")
		go helper.KafkaUnFollow(ctx, userIDObj.Hex(), userToUnFollowID, prod)
		go helper.KafkaFollowLog(ctx, userIDObj.Hex(), userToUnFollowID, "unfollowed", prod)
	}
}

func FollowRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		environement := c.GetString("env")
		var prod bool
		if environement == "prod" {
			prod = true
		} else {
			prod = false
		}

		// Get user ID from context
		userID, exists := c.Get("id")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
			return
		}

		userIDObj, ok := userID.(primitive.ObjectID)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		userToFollowID := c.Query("user_id")
		userToFollowIDObj, err := primitive.ObjectIDFromHex(userToFollowID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID to follow"})
			return
		}

		if userIDObj == userToFollowIDObj {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot follow yourself"})
			return
		}

		ctx := context.Background() // Consider using request-scoped context

		// Increment the 'followers' count of the user being followed
		update := bson.M{"$inc": bson.M{"follow_requests": 1}}
		_, err = UsersCollection(environement).UpdateOne(ctx, bson.M{"_id": userToFollowIDObj}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user to follow"})
			return
		}

		c.JSON(http.StatusOK, "REQUESTED")
		go helper.KafkaFollowRequest(ctx, userIDObj.Hex(), userToFollowID, prod)
		go helper.KafkaFollowLog(ctx, userIDObj.Hex(), userToFollowID, "requested", prod)
	}
}

func AcceptRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		environement := c.GetString("env")
		var prod bool
		if environement == "prod" {
			prod = true
		} else {
			prod = false
		}
    
		// Get user ID from context
		userID, exists := c.Get("id")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
			return
		}

		userIDObj, ok := userID.(primitive.ObjectID)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		userToFollowID := c.Query("user_id")
		userToFollowIDObj, err := primitive.ObjectIDFromHex(userToFollowID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID to follow"})
			return
		}

		ctx := context.Background() // Consider using request-scoped context

		// Increment the 'following' count of the user
		update := bson.M{"$inc": bson.M{"following": 1}}
		_, err = UsersCollection(environement).UpdateOne(ctx, bson.M{"_id": userToFollowIDObj}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		// Increment the 'followers' count of the user being followed
		update = bson.M{"$inc": bson.M{"followers": 1}}
		_, err = UsersCollection(environement).UpdateOne(ctx, bson.M{"_id": userIDObj}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user to follow"})
			return
		}

		// Increment the 'followers' count of the user being followed
		update = bson.M{"$inc": bson.M{"follow_requests": -1}}
		_, err = UsersCollection(environement).UpdateOne(ctx, bson.M{"_id": userIDObj}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user to follow"})
			return
		}

		c.JSON(http.StatusOK, "FOLLOWING")
		go helper.KafkaAcceptFollowRequest(ctx, userToFollowID, userIDObj.Hex(), prod)
		go helper.KafkaFollowLog(ctx, userIDObj.Hex(), userToFollowID, "accepted", prod)
	}
}

func DeclineRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		environement := c.GetString("env")
		var prod bool
		if environement == "prod" {
			prod = true
		} else {
			prod = false
		}

		// Get user ID from context
		userID, exists := c.Get("id")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
			return
		}

		userIDObj, ok := userID.(primitive.ObjectID)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		userToFollowID := c.Query("user_id")
		// userToFollowIDObj, err := primitive.ObjectIDFromHex(userToFollowID)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID to follow"})
		// 	return
		// }

		ctx := context.Background() // Consider using request-scoped context

		// Increment the 'followers' count of the user being followed
		update := bson.M{"$inc": bson.M{"follow_requests": -1}}
		_, err := UsersCollection(environement).UpdateOne(ctx, bson.M{"_id": userIDObj}, update)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user to follow"})
			return
		}

		c.JSON(http.StatusOK, "FOLLOW")

		go helper.KafkaDeclineFollowRequest(ctx, userToFollowID, userIDObj.Hex(), prod)
		go helper.KafkaFollowLog(ctx, userIDObj.Hex(), userToFollowID, "declined", prod)
	}
}

func CancelRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		environement := c.GetString("env")
		var prod bool
		if environement == "prod" {
			prod = true
		} else {
			prod = false
		}

		// Get user ID from context
		userID, exists := c.Get("id")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
			return
		}

		userIDObj, ok := userID.(primitive.ObjectID)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		userToFollowID := c.Query("user_id")
		userToFollowIDObj, err := primitive.ObjectIDFromHex(userToFollowID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID to follow"})
			return
		}

		ctx := context.Background() // Consider using request-scoped context

		// Increment the 'followers' count of the user being followed
		update := bson.M{"$inc": bson.M{"follow_requests": -1}}
		_, err = UsersCollection(environement).UpdateOne(ctx, bson.M{"_id": userToFollowIDObj}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user to follow"})
			return
		}

		c.JSON(http.StatusOK, "FOLLOW")

		go helper.KafkaCancelFollowRequest(ctx, userIDObj.Hex(), userToFollowID, prod)
		go helper.KafkaFollowLog(ctx, userIDObj.Hex(), userToFollowID, "cancelled", prod)
	}
}
