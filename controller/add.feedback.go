package controller

import (
	"log"
	"net/http"
	"time"
	"userrelation/database"
	"userrelation/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var FeedbackCollection *mongo.Collection = database.OrdersData(database.Orders, "Feedbacks")
var LocationCollection *mongo.Collection = database.RestaurantsData(database.Restaurants, "Location")

func AddFeedback() gin.HandlerFunc {
	return func(c *gin.Context) {
		restaurantID := c.Query("restaurant_id")
		locationID := c.Query("location_id")
		reservationID := c.Query("reservation_id")
		userID, exists := c.Get("id")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
			return
		}

		userIDObj := userID.(primitive.ObjectID)

		restaurantIDObj, err := primitive.ObjectIDFromHex(restaurantID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid restaurant ID"})
			return
		}

		locationIDObj, err := primitive.ObjectIDFromHex(locationID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid location ID"})
			return
		}

		reservationIDObj, err := primitive.ObjectIDFromHex(reservationID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reservation ID"})
			return
		}

		var feedback model.Feedback
		if err := c.BindJSON(&feedback); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		if feedback.Rating < 1 || feedback.Rating > 5 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Rating should be between 1 and 5"})
			return
		}

		// Ensure feedback.Rating is a float32
		feedbackRating := float32(feedback.Rating)

		// Populate feedback object with IDs and timestamp
		feedback.ID = primitive.NewObjectID()
		feedback.User_ID = userIDObj
		feedback.Restaurant_ID = restaurantIDObj
		feedback.Location_ID = locationIDObj
		feedback.Reservation_ID = reservationIDObj
		feedback.Created_At = time.Now()

		// Insert feedback into the FeedbackCollection
		_, err = FeedbackCollection.InsertOne(c, feedback)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save feedback"})
			return
		}

		// Update restaurant's rating and total reviews atomically
		restaurantPipeline := mongo.Pipeline{
			{{Key: "$set", Value: bson.D{
				{Key: "total_reviews", Value: bson.D{
					{Key: "$add", Value: bson.A{
						bson.D{{Key: "$ifNull", Value: bson.A{"$total_reviews", 0}}},
						1,
					}},
				}},
			}}},
			{{Key: "$set", Value: bson.D{
				{Key: "rating", Value: bson.D{
					{Key: "$add", Value: bson.A{
						bson.D{{Key: "$ifNull", Value: bson.A{bson.D{{Key: "$toDouble", Value: "$rating"}}, 0.0}}},
						bson.D{{Key: "$divide", Value: bson.A{
							bson.D{{Key: "$subtract", Value: bson.A{
								feedbackRating,
								bson.D{{Key: "$ifNull", Value: bson.A{bson.D{{Key: "$toDouble", Value: "$rating"}}, 0.0}}},
							}}},
							bson.D{{Key: "$toDouble", Value: "$total_reviews"}},
						}}},
					}},
				}},
				{Key: "updated_at", Value: time.Now()},
			}}},
		}

		// Update restaurant document
		_, err = RestaurantCollection.UpdateByID(c, restaurantIDObj, restaurantPipeline)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating restaurant rating"})
			return
		}

		// Update location's rating and total reviews atomically
		locationPipeline := mongo.Pipeline{
			{{Key: "$set", Value: bson.D{
				{Key: "total_reviews", Value: bson.D{
					{Key: "$add", Value: bson.A{
						bson.D{{Key: "$ifNull", Value: bson.A{"$total_reviews", 0}}},
						1,
					}},
				}},
			}}},
			{{Key: "$set", Value: bson.D{
				{Key: "rating", Value: bson.D{
					{Key: "$add", Value: bson.A{
						bson.D{{Key: "$ifNull", Value: bson.A{bson.D{{Key: "$toDouble", Value: "$rating"}}, 0.0}}},
						bson.D{{Key: "$divide", Value: bson.A{
							bson.D{{Key: "$subtract", Value: bson.A{
								feedbackRating,
								bson.D{{Key: "$ifNull", Value: bson.A{bson.D{{Key: "$toDouble", Value: "$rating"}}, 0.0}}},
							}}},
							bson.D{{Key: "$toDouble", Value: "$total_reviews"}},
						}}},
					}},
				}},
				{Key: "updated_at", Value: time.Now()},
			}}},
		}

		// Update location document
		_, err = LocationCollection.UpdateByID(c, locationIDObj, locationPipeline)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating location rating"})
			return
		}

		// Return success response
		c.JSON(http.StatusOK, gin.H{"message": "Feedback submitted and ratings updated successfully"})
	}
}

func CheckIfFeedback() gin.HandlerFunc {
	return func(c *gin.Context) {
		reservationID := c.Query("reservation_id")
		reservationIDObj, err := primitive.ObjectIDFromHex(reservationID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reservation ID"})
			return
		}

		filter := bson.M{"reservation_id": reservationIDObj}
		err = FeedbackCollection.FindOne(c, filter).Err() // Just checking for the error without decoding

		if err != nil {
			if err == mongo.ErrNoDocuments {
				// No feedback found for this reservation
				c.JSON(http.StatusNotFound, gin.H{"message": "No feedback found"})
				return
			}
			// Handle any other database errors
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking feedback"})
			return
		}

		// Feedback exists for this reservation
		c.JSON(http.StatusOK, gin.H{"message": "Feedback found"})
	}
}
