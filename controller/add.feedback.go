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

		// Validate rating value
		if feedback.Rating < 1 || feedback.Rating > 5 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Rating should be between 1 and 5"})
			return
		}

		// Populate feedback object with IDs and timestamp
		feedback.ID = primitive.NewObjectID()
		feedback.User_ID = userIDObj
		feedback.Restaurant_ID = restaurantIDObj
		feedback.Location_ID = locationIDObj
		feedback.Reservation_ID = reservationIDObj
		feedback.Created_At = time.Now()

		// Insert feedback into the orderCollection
		_, err = FeedbackCollection.InsertOne(c, feedback)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save feedback"})
			return
		}

		// Return success response
		c.JSON(http.StatusOK, "Feedback submitted successfully")
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
