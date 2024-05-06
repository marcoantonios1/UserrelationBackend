package controller

import (
	"context"
	"net/http"
	"userrelation/database"
	"userrelation/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UsersCollection *mongo.Collection = database.UsersData(database.Users, "Users")

func Follow() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from context
		userID, _ := c.Get("id")
		userIDObj, ok := userID.(primitive.ObjectID)
		if !ok {
			// Handle error if user ID is not a valid ObjectID
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		// Get user ID of the user to follow
		userToFollowID := c.Query("user_id")
		userToFollowIDObj, err := primitive.ObjectIDFromHex(userToFollowID)
		if err != nil {
			// Handle error if provided user ID is not a valid ObjectID
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		// Get the user from the database
		var followingUser model.User
		followingfilter := bson.M{"_id": userToFollowIDObj}
		err = UsersCollection.FindOne(context.Background(), followingfilter).Decode(&followingUser)
		if err != nil {
			// Handle error if user is not found
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		var User model.User
		filter := bson.M{"_id": userIDObj}
		err = UsersCollection.FindOne(context.Background(), filter).Decode(&User)
		if err != nil {
			// Handle error if user is not found
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		if followingUser.Private {
			c.JSON(http.StatusOK, "Requested")
		} else {
			User.Following++
			followingUser.Followers++
			update := bson.M{"$set": bson.M{"following": User.Following}}
			_, err = UsersCollection.UpdateOne(context.Background(), filter, update)
			if err != nil {
				// Handle error if update fails
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
				return
			}

			update = bson.M{"$set": bson.M{"followers": followingUser.Followers}}
			_, err = UsersCollection.UpdateOne(context.Background(), followingfilter, update)
			if err != nil {
				// Handle error if update fails
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
				return
			}

			// Return success message
			c.JSON(http.StatusOK, "Followed")
		}
	}
}

func UnFollow() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from context
		userID, _ := c.Get("id")
		userIDObj, ok := userID.(primitive.ObjectID)
		if !ok {
			// Handle error if user ID is not a valid ObjectID
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		// Get user ID of the user to follow
		userToUnFollowID := c.Query("user_id")
		userToUnFollowIDObj, err := primitive.ObjectIDFromHex(userToUnFollowID)
		if err != nil {
			// Handle error if provided user ID is not a valid ObjectID
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		// Get the user from the database
		var UnfollowingUser model.User
		Unfollowingfilter := bson.M{"_id": userToUnFollowIDObj}
		err = UsersCollection.FindOne(context.Background(), Unfollowingfilter).Decode(&UnfollowingUser)
		if err != nil {
			// Handle error if user is not found
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		var User model.User
		filter := bson.M{"_id": userIDObj}
		err = UsersCollection.FindOne(context.Background(), filter).Decode(&User)
		if err != nil {
			// Handle error if user is not found
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		User.Following--
		update := bson.M{"$set": bson.M{"following": User.Following}}
		_, err = UsersCollection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			// Handle error if update fails
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		UnfollowingUser.Followers--
		update = bson.M{"$set": bson.M{"followers": UnfollowingUser.Followers}}
		_, err = UsersCollection.UpdateOne(context.Background(), Unfollowingfilter, update)
		if err != nil {
			// Handle error if update fails
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		// Return success message
		c.JSON(http.StatusOK, "Followed")

	}
}
