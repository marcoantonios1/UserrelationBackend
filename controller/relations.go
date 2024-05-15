package controller

import (
	"context"
	"errors"
	"log"
	"net/http"
	"userrelation/model"

	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

func CheckSearchedUsersRelationship() gin.HandlerFunc {
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
						"userID":         userToFollowID,
						"userToFollowID": userIDObj.Hex(),
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

func RequestFollow() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		// Create a new driver for Neo4j
		driver, err := neo4j.NewDriverWithContext("neo4j://localhost:7687", neo4j.BasicAuth("neo4j", "12345678", ""))
		if err != nil {
			log.Fatal(err)
		}
		defer driver.Close(context.Background())

		// Create a new session
		session := driver.NewSession(context.Background(), neo4j.SessionConfig{DatabaseName: "usersRelations"})
		defer session.Close(context.Background())

		// Run the query to find users with REQUESTED relationship
		result, err := session.ExecuteRead(context.Background(),
			func(tx neo4j.ManagedTransaction) (interface{}, error) {
				result, err := tx.Run(context.Background(), `
                    MATCH (u:User)-[:REQUESTED]->(r:User {id: $userId})
                    RETURN u { .id, .username, .name, .image, .bio, .private } AS user
                `,
					map[string]interface{}{
						"userId": userIDObj.Hex(),
					},
				)
				if err != nil {
					return nil, err
				}

				var users []model.Neo4jUser
				for result.NextRecord(context.Background(), nil) {
					userNode, ok := result.Record().Get("user")
					if !ok {
						return nil, errors.New("failed to get user node")
					}
					userMap := userNode.(map[string]interface{})

					user := model.Neo4jUser{
						ID:        userMap["id"].(string),
						UserName:  userMap["username"].(string),
						Name:      userMap["name"].(string),
						Image:     userMap["image"].(string),
						Biography: userMap["bio"].(string),
						Private:   userMap["private"].(bool),
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

		// Return the list of users with matching structure
		c.JSON(http.StatusOK, result)
	}
}
