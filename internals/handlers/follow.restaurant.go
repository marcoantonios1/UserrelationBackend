package handlers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"userrelation/internals/models"
	helper "userrelation/internals/utils"

	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UnFollowRestaurant() gin.HandlerFunc {
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
		_, err = UsersCollection(environement).UpdateOne(ctx, bson.M{"_id": userIDObj}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		// Decrement the 'followers' count of the user being unfollowed
		update = bson.M{"$inc": bson.M{"followers": -1}}
		_, err = RestaurantCollection(environement).UpdateOne(ctx, bson.M{"_id": userToUnFollowIDObj}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user to unfollow"})
			return
		}

		c.JSON(http.StatusOK, "FOLLOW")
		go helper.KafkaUnFollowRestaurant(ctx, userIDObj.Hex(), userToUnFollowID, prod)
		go helper.KafkaRestaurantFollowLog(ctx, userIDObj.Hex(), userToUnFollowID, "unfollowed",prod)
	}
}

func FollowRestaurant() gin.HandlerFunc {
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
		_, err = UsersCollection(environement).UpdateOne(ctx, bson.M{"_id": userIDObj}, update)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		// Increment the 'followers' count of the user being followed
		update = bson.M{"$inc": bson.M{"followers": 1}}
		_, err = RestaurantCollection(environement).UpdateOne(ctx, bson.M{"_id": userToFollowIDObj}, update)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user to follow"})
			return
		}

		c.JSON(http.StatusOK, "FOLLOWING")
		go helper.KafkaFollowRestaurant(ctx, userIDObj.Hex(), userToFollowID, prod)
		go helper.KafkaRestaurantFollowLog(ctx, userIDObj.Hex(), userToFollowID, "followed",prod)
	}
}

func ViewFollowedRestaurant() gin.HandlerFunc {
	return func(c *gin.Context) {
		environement := c.GetString("env")
		// var prod bool
		// if environement == "prod" {
		// 	prod = true
		// } else {
		// 	prod = false
		// }

		usersearchID := c.Query("id")
		userID, exists := c.Get("id")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
			return
		}

		if usersearchID == "" {
			userIDObj, ok := userID.(primitive.ObjectID)
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
				return
			}
			usersearchID = userIDObj.Hex()
		}
		// Create a new driver for Neo4j
		driver, err := neo4j.NewDriverWithContext(Neo4j(environement), neo4j.BasicAuth(Neo4j_User, Neo4j_Password(environement), ""))

		if err != nil {
			log.Fatal(err)
		}
		defer driver.Close(context.Background())

		// Create a new session
		session := driver.NewSession(context.Background(), neo4j.SessionConfig{DatabaseName: Neo4j_Database(environement)})
		defer session.Close(context.Background())

		// Run the query to find users with REQUESTED relationship
		result, err := session.ExecuteRead(context.Background(),
			func(tx neo4j.ManagedTransaction) (interface{}, error) {
				result, err := tx.Run(context.Background(), `
                    MATCH (r:User {id: $userId})-[:FOLLOWING]->(u:Restaurant)
                    RETURN u { .id,  .name, .image, .description } AS user
                `,
					map[string]interface{}{
						"userId": usersearchID,
					},
				)
				if err != nil {
					return nil, err
				}

				var users []models.Restaurant
				for result.NextRecord(context.Background(), nil) {
					userNode, ok := result.Record().Get("user")
					if !ok {
						return nil, errors.New("failed to get user node")
					}
					userMap := userNode.(map[string]interface{})

					image, ok := userMap["image"].(string)
					if !ok {
						image = "" // or any default value
					}
					name, ok := userMap["name"].(string)
					if !ok {
						name = "" // or any default value
					}
					description, ok := userMap["description"].(string)
					if !ok {
						description = "" // or any default value
					}
					user := models.Restaurant{
						Restaurant_ID:   userMap["id"].(string),
						Restaurant_Name: name,
						Description:     description,
						Image:           image,
					}

					users = append(users, user)
				}

				return users, nil
			},
		)

		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if len(result.([]models.Restaurant)) == 0 {
			c.JSON(http.StatusOK, []models.Restaurant{})
			return
		}

		// Return the list of users with matching structure
		c.JSON(http.StatusOK, result)
	}
}
