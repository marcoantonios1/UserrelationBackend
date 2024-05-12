package controller

import (
	"context"
	"errors"
	"log"
	"net/http"
	"userrelation/database"
	"userrelation/helper"

	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UsersCollection *mongo.Collection = database.UsersData(database.Users, "Users")

func CheckUsersRelationship() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("id")
		userToFollowID := c.Query("user_id")

		// Convert userID to string
		userIDObj := userID.(primitive.ObjectID)

		// Create a new driver for Neo4j
		driver, err := neo4j.NewDriverWithContext("neo4j://localhost:7687", neo4j.BasicAuth("neo4j", "12345678", ""))
		if err != nil {
			log.Fatal(err)
		}
		defer driver.Close(context.Background())

		// Create a new session
		session := driver.NewSession(context.Background(), neo4j.SessionConfig{DatabaseName: "usersRelations"})
		defer session.Close(context.Background())

		// Run the query to check for relationship
		result, err := session.ExecuteRead(context.Background(),
			func(tx neo4j.ManagedTransaction) (any, error) {
				result, err := tx.Run(context.Background(), "MATCH (a:User)-[r]->(b:User) WHERE a.id = $userID AND b.id = $userToFollowID RETURN type(r)",
					map[string]interface{}{
						"userID":         userIDObj.Hex(),
						"userToFollowID": userToFollowID,
					})
				if err != nil {
					return nil, err // Handle error here (e.g., return specific error)
				}

				records, err := result.Collect(context.Background())
				if err != nil {
					return nil, err // Handle error here
				}

				// Check if any records exist (meaning a relationship exists)
				if len(records) == 0 {
					return "FOLLOW", nil // Indicate no relationship found
				}

				// Extract the relationship type from the first record
				value, ok := records[0].Get("type(r)")
				if !ok {
					// Handle case where key "type(r)" doesn't exist (log or return error)
					log.Println("Error: Key 'type(r)' not found in record")
					return nil, errors.New("unexpected record structure") // Replace with appropriate error
				}

				relationshipType := value.(string)
				return relationshipType, nil
			})
		if err != nil {
			log.Fatal(err)
		}

		// Respond based on the result
		relation, ok := result.(string)
		if !ok {
			// Handle unexpected result type (log or return error)
			log.Printf("Unexpected result type: %T", result)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": relation,
		})
	}
}

func CheckRestaurantRelationship() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("id")
		restaurantToFollowID := c.Query("resto_id")

		ueserIdObj := userID.(primitive.ObjectID)

		// Create a new driver for Neo4j
		driver, err := neo4j.NewDriverWithContext("neo4j://localhost:7687", neo4j.BasicAuth("neo4j", "12345678", ""))
		if err != nil {
			log.Fatal(err)
		}
		defer driver.Close(context.Background())

		// Create a new session
		session := driver.NewSession(context.Background(), neo4j.SessionConfig{DatabaseName: "usersRelations"})
		defer session.Close(context.Background())

		// Run the query to check if the user is following the other user
		result, err := session.ExecuteRead(context.Background(),
			func(tx neo4j.ManagedTransaction) (any, error) {
				result, err := tx.Run(context.Background(), "MATCH (a:User)-[r]->(b:Restaurant ) WHERE a.id = $userID AND b.id = $restaurantToFollowID RETURN type(r)",
					map[string]interface{}{
						"userID":               ueserIdObj.Hex(),
						"restaurantToFollowID": restaurantToFollowID,
					})
				if err != nil {
					return nil, err // Handle error here (e.g., return specific error)
				}

				records, err := result.Collect(context.Background())
				if err != nil {
					return nil, err // Handle error here
				}

				// Check if any records exist (meaning a relationship exists)
				if len(records) == 0 {
					return "FOLLOW", nil // Indicate no relationship found
				}

				// Extract the relationship type from the first record
				value, ok := records[0].Get("type(r)")
				if !ok {
					// Handle case where key "type(r)" doesn't exist (log or return error)
					log.Println("Error: Key 'type(r)' not found in record")
					return nil, errors.New("unexpected record structure") // Replace with appropriate error
				}

				relationshipType := value.(string)
				return relationshipType, nil
			})
		if err != nil {
			log.Fatal(err)
		}

		// Respond based on the result
		relation, ok := result.(string)
		if !ok {
			// Handle unexpected result type (log or return error)
			log.Printf("Unexpected result type: %T", result)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": relation,
		})
	}
}

func CountRestaurantFollowers() gin.HandlerFunc {
	return func(c *gin.Context) {
		restaurantID := c.Query("resto_id")

		// Create a new driver for Neo4j
		driver, err := neo4j.NewDriverWithContext("neo4j://localhost:7687", neo4j.BasicAuth("neo4j", "12345678", ""))
		if err != nil {
			log.Fatal(err)
		}
		defer driver.Close(context.Background())

		// Create a new session
		session := driver.NewSession(context.Background(), neo4j.SessionConfig{DatabaseName: "usersRelations"})
		defer session.Close(context.Background())

		// Run the query to check if the user is following the other user
		result, err := session.ExecuteRead(context.Background(),
			func(tx neo4j.ManagedTransaction) (any, error) {
				result, err := tx.Run(context.Background(), "MATCH (u:User)-[:FOLLOWS]->(r:Restaurant) WHERE r.id = $restaurantID RETURN count(u) AS followerCount",
					map[string]interface{}{
						"restaurantID": restaurantID,
					},
				)
				if err != nil {
					return nil, err
				}

				// Check for no followers
				if !result.NextRecord(context.Background(), nil) {
					return 0, nil // Return 0 if no followers found
				}

				// Get the follower count
				value, ok := result.Record().Get("followerCount")
				if !ok {
					return nil, errors.New("failed to get follower count")
				}

				followerCount, ok := value.(int64)
				if !ok {
					return nil, errors.New("unexpected type for follower count")
				}

				return followerCount, nil
			},
		)

		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the follower count
		c.JSON(http.StatusOK, gin.H{"result": result.(int64)})
	}
}

func Follow() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		_, err = UsersCollection.UpdateOne(ctx, bson.M{"_id": userIDObj}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		// Increment the 'followers' count of the user being followed
		update = bson.M{"$inc": bson.M{"followers": 1}}
		_, err = UsersCollection.UpdateOne(ctx, bson.M{"_id": userToFollowIDObj}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user to follow"})
			return
		}

		c.JSON(http.StatusOK, "Followed")
		go helper.KafkaFollow(ctx, userIDObj.Hex(), userToFollowID)
	}
}

func UnFollow() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		_, err = UsersCollection.UpdateOne(ctx, bson.M{"_id": userIDObj}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		// Decrement the 'followers' count of the user being unfollowed
		update = bson.M{"$inc": bson.M{"followers": -1}}
		_, err = UsersCollection.UpdateOne(ctx, bson.M{"_id": userToUnFollowIDObj}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user to unfollow"})
			return
		}

		c.JSON(http.StatusOK, "Unfollowed")
		go helper.KafkaUnFollow(ctx, userIDObj.Hex(), userToUnFollowID)
	}
}

func FollowRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		update := bson.M{"$inc": bson.M{"follow_requests": 1}}
		_, err = UsersCollection.UpdateOne(ctx, bson.M{"_id": userToFollowIDObj}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user to follow"})
			return
		}

		c.JSON(http.StatusOK, "Requested")
		go helper.KafkaFollowRequest(ctx, userIDObj.Hex(), userToFollowID)
	}
}

func AcceptRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		_, err = UsersCollection.UpdateOne(ctx, bson.M{"_id": userToFollowIDObj}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		// Increment the 'followers' count of the user being followed
		update = bson.M{"$inc": bson.M{"followers": 1}}
		_, err = UsersCollection.UpdateOne(ctx, bson.M{"_id": userIDObj}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user to follow"})
			return
		}

		c.JSON(http.StatusOK, "Followed")
		go helper.KafkaAcceptFollowRequest(ctx, userIDObj.Hex(), userToFollowID)
	}
}

func DeclineRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		_, err = UsersCollection.UpdateOne(ctx, bson.M{"_id": userToFollowIDObj}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user to follow"})
			return
		}

		c.JSON(http.StatusOK, "Followed")

		go helper.KafkaDeclineFollowRequest(ctx, userIDObj.Hex(), userToFollowID)
	}
}
