package handlers

// import (
// 	"context"
// 	"log"
// 	"net/http"
// 	"userrelation/model"

// 	"github.com/gin-gonic/gin"
// 	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// func GetOtherUserAndRelationship() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Get user IDs from context and query
// 		userID, exists := c.Get("id")
// 		if !exists {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
// 			return
// 		}

// 		otherUserID := c.Query("user_id")

// 		userIdObj := userID.(primitive.ObjectID)

// 		// Create a new driver for Neo4j
// 		driver, err := neo4j.NewDriverWithContext("neo4j://localhost:7687", neo4j.BasicAuth("neo4j", "12345678", ""))
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer driver.Close(context.Background())

// 		// Create a new session
// 		session := driver.NewSession(context.Background(), neo4j.SessionConfig{DatabaseName: "usersRelations"})
// 		defer session.Close(context.Background())

// 		// Run the query to check for relationship
// 		result, err := session.ExecuteRead(context.Background(),
// 			func(tx neo4j.ManagedTransaction) (any, error) {
// 				result, err := tx.Run(context.Background(), "MATCH (a:User)-[r]->(b:User) WHERE a.id = $userID AND b.id = $otherUserID RETURN b, type(r)", map[string]interface{}{
// 					"userID":      userIdObj.Hex(),
// 					"otherUserID": otherUserID,
// 				})
// 				if err != nil {
// 					return nil, err
// 				}

// 				// Initialize userInfo and relationshipType
// 				var userInfo *model.Neo4jUser
// 				relationshipType := "FOLLOW" // Default relation

// 				// Get the first record
// 				record, err := result.Single(context.Background())
// 				if err != nil {
// 					return nil, err
// 				}

// 				if record != nil {
// 					userNode := record.Values[0].(neo4j.Node)
// 					if len(record.Values) > 1 {
// 						relationshipType = record.Values[1].(string)
// 					}

// 					// Extract user info from the node properties
// 					userInfo = &model.Neo4jUser{
// 						ID:        userNode.Props["id"].(string),
// 						UserName:  userNode.Props["username"].(*string),
// 						Name:      userNode.Props["name"].(*string),
// 						Image:     userNode.Props["image"].(*string),
// 						Biography: userNode.Props["bio"].(*string),
// 						Followers: userNode.Props["followers"].(uint32),
// 						Following: userNode.Props["following"].(uint32),
// 						Private:   userNode.Props["private"].(bool),
// 						Relation:  relationshipType,
// 					}
// 				}

// 				return userInfo, nil
// 			})
// 		if err != nil {
// 			log.Print(err)
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get other user info and relationship"})
// 			return
// 		}

// 		// Return the other user info and relationship type
// 		c.JSON(http.StatusOK, result)
// 	}
// }
