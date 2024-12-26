package handlers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"userrelation/internals/models"

	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckUsersRelationship() gin.HandlerFunc {
	return func(c *gin.Context) {
		environement := c.GetString("env")
		// var prod bool
		// if environement == "prod" {
		// 	prod = true
		// } else {
		// 	prod = false
		// }
		userID, _ := c.Get("id")
		userToFollowID := c.Query("user_id")

		// Convert userID to string
		userIDObj := userID.(primitive.ObjectID)

		// Create a new driver for Neo4j
		driver, err := neo4j.NewDriverWithContext(Neo4j(environement), neo4j.BasicAuth(Neo4j_User, Neo4j_Password(environement), ""))
		if err != nil {
			log.Fatal(err)
		}
		defer driver.Close(context.Background())

		// Create a new session
		session := driver.NewSession(context.Background(), neo4j.SessionConfig{DatabaseName: Neo4j_Database(environement)})
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
		environement := c.GetString("env")
		// var prod bool
		// if environement == "prod" {
		// 	prod = true
		// } else {
		// 	prod = false
		// }
		userID, _ := c.Get("id")
		userToFollowID := c.Query("user_id")

		// Convert userID to string
		userIDObj := userID.(primitive.ObjectID)

		// Create a new driver for Neo4j
		driver, err := neo4j.NewDriverWithContext(Neo4j(environement), neo4j.BasicAuth(Neo4j_User, Neo4j_Password(environement), ""))
		if err != nil {
			log.Fatal(err)
		}
		defer driver.Close(context.Background())

		// Create a new session
		session := driver.NewSession(context.Background(), neo4j.SessionConfig{DatabaseName: Neo4j_Database(environement)})
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
		environement := c.GetString("env")
		// var prod bool
		// if environement == "prod" {
		// 	prod = true
		// } else {
		// 	prod = false
		// }
		userID, _ := c.Get("id")
		restaurantToFollowID := c.Query("resto_id")

		userIdObj := userID.(primitive.ObjectID)

		// Create a new driver for Neo4j
		driver, err := neo4j.NewDriverWithContext(Neo4j(environement), neo4j.BasicAuth(Neo4j_User, Neo4j_Password(environement), ""))
		if err != nil {
			log.Fatal(err)
		}
		defer driver.Close(context.Background())

		// Create a new session
		session := driver.NewSession(context.Background(), neo4j.SessionConfig{DatabaseName: Neo4j_Database(environement)})
		defer session.Close(context.Background())

		// Run the query to check if the user has a FOLLOWING relationship with the restaurant
		result, err := session.ExecuteRead(context.Background(),
			func(tx neo4j.ManagedTransaction) (any, error) {
				query := `
					MATCH (a:User)-[r:FOLLOWING]->(b:Restaurant)
					WHERE a.id = $userID AND b.id = $restaurantToFollowID
					RETURN COUNT(r) AS followingCount
				`
				params := map[string]interface{}{
					"userID":               userIdObj.Hex(),
					"restaurantToFollowID": restaurantToFollowID,
				}

				res, err := tx.Run(context.Background(), query, params)
				if err != nil {
					return nil, err
				}

				record, err := res.Single(context.Background())
				if err != nil {
					// No "FOLLOWING" relationship was found
					return "FOLLOW", nil
				}

				// Get the count of FOLLOWING relationships
				followingCount, _ := record.Get("followingCount")
				count, ok := followingCount.(int64)
				if !ok {
					return nil, errors.New("unexpected count type")
				}

				// If there is a FOLLOWING relationship, return "FOLLOWING"
				if count > 0 {
					return "FOLLOWING", nil
				}
				return "FOLLOW", nil
			})
		if err != nil {
			log.Fatal(err)
		}

		// Respond based on the result
		relation, ok := result.(string)
		if !ok {
			log.Printf("Unexpected result type: %T", result)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": relation,
		})
	}
}

func RequestFollow() gin.HandlerFunc {
	return func(c *gin.Context) {
		environement := c.GetString("env")
		// var prod bool
		// if environement == "prod" {
		// 	prod = true
		// } else {
		// 	prod = false
		// }
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

				var users []models.Neo4jUser
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
					bio, ok := userMap["bio"].(string)
					if !ok {
						bio = "" // or any default value
					}
					private, ok := userMap["private"].(bool)
					if !ok {
						private = false // or any default value
					}
					user := models.Neo4jUser{
						ID:        userMap["id"].(string),
						UserName:  userMap["username"].(string),
						Name:      name,
						Image:     image,
						Biography: bio,
						Private:   private,
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

		if len(result.([]models.Neo4jUser)) == 0 {
			c.JSON(http.StatusOK, []models.Neo4jUser{})
			return
		}

		// Return the list of users with matching structure
		c.JSON(http.StatusOK, result)
	}
}

func ViewFollowing() gin.HandlerFunc {
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
                    MATCH (r:User {id: $userId})-[:FOLLOWING]->(u:User)
                    RETURN u { .id, .username, .name, .image, .bio, .private } AS user
                `,
					map[string]interface{}{
						"userId": usersearchID,
					},
				)
				if err != nil {
					return nil, err
				}

				var users []models.Neo4jUser
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
					bio, ok := userMap["bio"].(string)
					if !ok {
						bio = "" // or any default value
					}
					private, ok := userMap["private"].(bool)
					if !ok {
						private = false // or any default value
					}
					user := models.Neo4jUser{
						ID:        userMap["id"].(string),
						UserName:  userMap["username"].(string),
						Name:      name,
						Image:     image,
						Biography: bio,
						Private:   private,
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

		if len(result.([]models.Neo4jUser)) == 0 {
			c.JSON(http.StatusOK, []models.Neo4jUser{})
			return
		}

		// Return the list of users with matching structure
		c.JSON(http.StatusOK, result)
	}
}

func ViewFollowers() gin.HandlerFunc {
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
                    MATCH (u:User)-[:FOLLOWING]->(r:User {id: $userId})
                    RETURN u { .id, .username, .name, .image, .bio, .private } AS user
                `,
					map[string]interface{}{
						"userId": usersearchID,
					},
				)
				if err != nil {
					return nil, err
				}

				var users []models.Neo4jUser
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
					bio, ok := userMap["bio"].(string)
					if !ok {
						bio = "" // or any default value
					}
					private, ok := userMap["private"].(bool)
					if !ok {
						private = false // or any default value
					}
					user := models.Neo4jUser{
						ID:        userMap["id"].(string),
						UserName:  userMap["username"].(string),
						Name:      name,
						Image:     image,
						Biography: bio,
						Private:   private,
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

		if len(result.([]models.Neo4jUser)) == 0 {
			c.JSON(http.StatusOK, []models.Neo4jUser{})
			return
		}

		// Return the list of users with matching structure
		c.JSON(http.StatusOK, result)
	}
}

func GetMutualFollowers() gin.HandlerFunc {
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

		userIDObj, ok := userID.(primitive.ObjectID)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		driver, err := neo4j.NewDriverWithContext(Neo4j(environement), neo4j.BasicAuth(Neo4j_User, Neo4j_Password(environement), ""))
		if err != nil {
			log.Fatal(err)
		}
		defer driver.Close(context.Background())

		// Create a new session
		session := driver.NewSession(context.Background(), neo4j.SessionConfig{DatabaseName: Neo4j_Password(environement)})
		defer session.Close(context.Background())

		result, err := session.ExecuteRead(context.Background(),
			func(tx neo4j.ManagedTransaction) (interface{}, error) {
				result, err := tx.Run(context.Background(), `
				MATCH (u1:User {id: $userID1})<-[:FOLLOWING]-(m:User)-[:FOLLOWING]->(u2:User {id: $userID2})
				RETURN m as mutualFollowers
                `,
					map[string]interface{}{
						"userID1": usersearchID,
						"userID2": userIDObj.Hex(),
					},
				)
				if err != nil {
					return nil, err
				}

				var users []models.Neo4jUser
				for result.NextRecord(context.Background(), nil) {
					userNode, ok := result.Record().Get("mutualFollowers")
					if !ok {
						return nil, errors.New("failed to get user node")
					}
					userNodeDb, ok := userNode.(dbtype.Node)
					if !ok {
						return nil, errors.New("failed to convert to dbtype.Node")
					}
					userMap := userNodeDb.Props

					image, ok := userMap["image"].(string)
					if !ok {
						image = "" // or any default value
					}
					name, ok := userMap["name"].(string)
					if !ok {
						name = "" // or any default value
					}
					bio, ok := userMap["bio"].(string)
					if !ok {
						bio = "" // or any default value
					}
					private, ok := userMap["private"].(bool)
					if !ok {
						private = false // or any default value
					}
					user := models.Neo4jUser{
						ID:        userMap["id"].(string),
						UserName:  userMap["username"].(string),
						Name:      name,
						Image:     image,
						Biography: bio,
						Private:   private,
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

		if len(result.([]models.Neo4jUser)) == 0 {
			c.JSON(http.StatusOK, []models.Neo4jUser{})
			return
		}

		// Return the list of users with matching structure
		c.JSON(http.StatusOK, result)
	}
}

func GetMutualFollowersCount() gin.HandlerFunc {
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

		userIDObj, ok := userID.(primitive.ObjectID)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		driver, err := neo4j.NewDriverWithContext(Neo4j(environement), neo4j.BasicAuth(Neo4j_User, Neo4j_Password(environement), ""))
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			if err := driver.Close(context.Background()); err != nil {
				log.Printf("Error closing driver: %v", err)
			}
		}()

		// Create a new session
		session := driver.NewSession(context.Background(), neo4j.SessionConfig{DatabaseName: Neo4j_Database(environement)})
		defer func() {
			if err := session.Close(context.Background()); err != nil {
				log.Printf("Error closing session: %v", err)
			}
		}()

		result, err := session.ExecuteRead(context.Background(),
			func(tx neo4j.ManagedTransaction) (interface{}, error) {
				query := `
                MATCH (u1:User {id: $userID1})<-[:FOLLOWING]-(m:User)-[:FOLLOWING]->(u2:User {id: $userID2})
                RETURN count(m) as mutualFollowersCount
                `
				params := map[string]interface{}{
					"userID1": usersearchID,
					"userID2": userIDObj.Hex(),
				}

				result, err := tx.Run(context.Background(), query, params)
				if err != nil {
					return nil, err
				}

				if result.Next(context.Background()) {
					record := result.Record()
					if count, ok := record.Get("mutualFollowersCount"); ok {
						return count.(int64), nil
					}
				}

				return int64(0), nil
			},
		)

		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the count of mutual followers
		c.JSON(http.StatusOK, gin.H{"mutualFollowersCount": result})
	}
}
