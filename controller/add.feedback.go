package controller

import (
	"context"
	"log"
	"net/http"
	"time"
	"userrelation/database"
	"userrelation/helper"
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

		// Check if feedback comment is empty
		increaseTotalReviews := false
		if feedback.Feedback != "" {
			increaseTotalReviews = true
		}

		// Populate feedback object with IDs and timestamp
		feedback.ID = primitive.NewObjectID()
		feedback.User_ID = userIDObj
		feedback.Restaurant_ID = restaurantIDObj
		feedback.Location_ID = locationIDObj
		feedback.Reservation_ID = reservationIDObj
		feedback.Created_At = time.Now()

		if feedback.Feedback == "" {
			_, err = FeedbackCollection.InsertOne(c, feedback)
			if err != nil {
				log.Print(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save feedback"})
				return
			}
		} else {
			go helper.KafkaLeaveFeedbackRestaurant(context.Background(), feedback.User_ID.Hex(), feedback.Restaurant_ID.Hex(), feedback.Location_ID.Hex(), feedback.Reservation_ID.Hex(), feedback.Rating, feedback.Feedback, feedback.Created_At.String())
			feedback.Feedback = ""
			_, err = FeedbackCollection.InsertOne(c, feedback)
			if err != nil {
				log.Print(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save feedback"})
				return
			}
		}

		// Update restaurant's rating, total_ratings, and total_reviews
		restaurantFilter := bson.M{"_id": restaurantIDObj}

		// Fetch current rating and total_ratings
		var restaurant model.RestaurantFeedbackUpdate
		err = RestaurantCollection.FindOne(c, restaurantFilter).Decode(&restaurant)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching restaurant data"})
			return
		}

		currentRating := restaurant.Rating
		currentTotalRatings := restaurant.TotalRatings

		// Compute new total_ratings and new rating
		newTotalRatings := currentTotalRatings + 1
		newRating := ((currentRating * float64(currentTotalRatings)) + float64(feedbackRating)) / float64(newTotalRatings)

		restaurantUpdate := bson.M{
			"$set": bson.M{
				"rating":     newRating,
				"updated_at": time.Now(),
			},
			"$inc": bson.M{
				"total_ratings": 1,
			},
		}

		if increaseTotalReviews {
			restaurantUpdate["$inc"].(bson.M)["total_reviews"] = 1
		}

		// Update restaurant document
		_, err = RestaurantCollection.UpdateOne(c, restaurantFilter, restaurantUpdate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating restaurant rating"})
			return
		}

		// Update location's rating, total_ratings, and total_reviews
		locationFilter := bson.M{"_id": locationIDObj}

		// Fetch current rating and total_ratings for location
		var location model.Location
		err = LocationCollection.FindOne(c, locationFilter).Decode(&location)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching location data"})
			return
		}

		currentLocationRating := location.Rating
		currentLocationTotalRatings := location.TotalRatings

		// Compute new total_ratings and new rating for location
		newLocationTotalRatings := currentLocationTotalRatings + 1
		newLocationRating := ((currentLocationRating * float64(currentLocationTotalRatings)) + float64(feedbackRating)) / float64(newLocationTotalRatings)

		locationUpdate := bson.M{
			"$set": bson.M{
				"rating":     newLocationRating,
				"updated_at": time.Now(),
			},
			"$inc": bson.M{
				"total_ratings": 1,
			},
		}

		if increaseTotalReviews {
			locationUpdate["$inc"].(bson.M)["total_reviews"] = 1
		}

		// Update location document
		_, err = LocationCollection.UpdateOne(c, locationFilter, locationUpdate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating location rating"})
			return
		}

		// --- New Code Ends Here ---

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
