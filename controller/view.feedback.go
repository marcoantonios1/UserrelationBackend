package controller

import (
	"context"
	"errors"
	"log"
	"net/http"
	"userrelation/model"

	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ViewRestaurantFeedback() gin.HandlerFunc {
	return func(c *gin.Context) {
		restaurantID := c.Query("restaurantId")
		locationID := c.Query("locationId")
		isLocation := false
		if restaurantID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Restaurant ID is required"})
			return
		}
		if locationID != "" {
			isLocation = true
		}

		// Create a new driver for Neo4j
		driver, err := neo4j.NewDriverWithContext(Neo4j, neo4j.BasicAuth(Neo4j_User, Neo4j_Password, ""))
		if err != nil {
			log.Fatal(err)
		}
		defer driver.Close(context.Background())

		// Create a new session
		session := driver.NewSession(context.Background(), neo4j.SessionConfig{DatabaseName: "usersRelations"})
		defer session.Close(context.Background())

		// Run the query to get feedback with user info
		result, err := session.ExecuteRead(context.Background(),
			func(tx neo4j.ManagedTransaction) (interface{}, error) {
				var result neo4j.ResultWithContext
				if isLocation {
					result, err = tx.Run(context.Background(), `
                    MATCH (u:User)-[r:REVIEWED]->(restaurant:Restaurant {id: $restaurantId})
                    WHERE r.locationId = $locationId
                    RETURN u { .id, .username, .image } AS user, 
                           r { .feedback, .rating, .createdAt } AS review
                `,
						map[string]interface{}{
							"restaurantId": restaurantID,
							"locationId":   locationID,
						},
					)
				} else {
					result, err = tx.Run(context.Background(), `
                    MATCH (u:User)-[r:REVIEWED]->(restaurant:Restaurant {id: $restaurantId})
                    RETURN u { .id, .username, .image } AS user, 
                           r { .feedback, .rating, .createdAt } AS review
                `,
						map[string]interface{}{
							"restaurantId": restaurantID,
						},
					)
				}
				if err != nil {
					return nil, err
				}

				var feedbacks []model.RestaurantFeedback
				for result.NextRecord(context.Background(), nil) {
					// Extract user and review details
					userNode, ok := result.Record().Get("user")
					if !ok {
						return nil, errors.New("failed to get user node")
					}
					reviewNode, ok := result.Record().Get("review")
					if !ok {
						return nil, errors.New("failed to get review node")
					}

					// Map user details
					userMap := userNode.(map[string]interface{})
					image, _ := userMap["image"].(string)
					username, _ := userMap["username"].(string)
					userID, _ := userMap["id"].(string)

					// Map review details
					reviewMap := reviewNode.(map[string]interface{})
					feedback, _ := reviewMap["feedback"].(string)
					rating64, _ := reviewMap["rating"].(int64)
					rating := int(rating64)
					createdAt, _ := reviewMap["createdAt"].(string)

					// Create a new feedback item
					feedbacks = append(feedbacks, model.RestaurantFeedback{
						UserID:    userID,
						Username:  username,
						Image:     image,
						Feedback:  feedback,
						Rating:    rating,
						CreatedAt: createdAt,
					})
				}

				return feedbacks, nil
			},
		)

		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		feedbacks, ok := result.([]model.RestaurantFeedback)
		if !ok {
			feedbacks = []model.RestaurantFeedback{}
		}

		// Return the list of feedback items
		c.JSON(http.StatusOK, feedbacks)
	}
}

func GetStarCounts() gin.HandlerFunc {
	return func(c *gin.Context) {
		restaurantID := c.Query("restaurantId")
		isLocation := false
		if restaurantID == "" {
			restaurantID = c.Query("locationId")
			isLocation = true
			if restaurantID == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Restaurant ID is required"})
				return
			}
		}
		restaurantIDObj, err := primitive.ObjectIDFromHex(restaurantID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid restaurant ID"})
			return
		}

		// MongoDB aggregation pipeline to count each star rating
		var pipeline mongo.Pipeline
		if isLocation {
			pipeline = mongo.Pipeline{
				{{Key: "$match", Value: bson.D{{Key: "location_id", Value: restaurantIDObj}}}},
				{{Key: "$group", Value: bson.D{
					{Key: "_id", Value: "$rating"},
					{Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}},
				}}},
			}
		} else {
			pipeline = mongo.Pipeline{
				{{Key: "$match", Value: bson.D{{Key: "restaurant_id", Value: restaurantIDObj}}}},
				{{Key: "$group", Value: bson.D{
					{Key: "_id", Value: "$rating"},
					{Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}},
				}}},
			}
		}

		// Execute the aggregation
		cursor, err := FeedbackCollection.Aggregate(context.Background(), pipeline)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer cursor.Close(context.Background())

		// Initialize star counts
		var counts model.TotalCountPreStar
		for cursor.Next(context.Background()) {
			var result struct {
				ID    int `bson:"_id"`
				Count int `bson:"count"`
			}
			if err := cursor.Decode(&result); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// Map counts to the struct based on star rating
			switch result.ID {
			case 5:
				counts.FiveStars = result.Count
			case 4:
				counts.FourStars = result.Count
			case 3:
				counts.ThreeStars = result.Count
			case 2:
				counts.TwoStars = result.Count
			case 1:
				counts.OneStar = result.Count
			}
		}

		if err := cursor.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the counts
		c.JSON(http.StatusOK, counts)
	}
}
