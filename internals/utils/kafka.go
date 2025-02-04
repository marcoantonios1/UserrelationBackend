package utils

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

var KafkaUrl string = os.Getenv("KAFKA_URL")

func KafkaFollow(ctx context.Context, followerId string, userToFollowID string, prod bool) {
	// Kafka configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{KafkaUrl},
		Topic:    "user_relation",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	// Prepare and send the message
	message := map[string]interface{}{
		"followerId": followerId,
		"followeeId": userToFollowID,
		"following":  true,
		"prod":       prod,
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Println("Error marshaling Kafka message:", err)
		return
	}

	msg := kafka.Message{
		Key:   []byte(followerId),
		Value: messageJSON,
	}

	err = kafkaWriter.WriteMessages(ctx, msg)
	if err != nil {
		// Handle Kafka sending error (log or return an error)
		log.Println("Error sending message to Kafka:", err)
	}
}

func KafkaUnFollow(ctx context.Context, followerId string, userToFollowID string, prod bool) {
	// Kafka configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{KafkaUrl},
		Topic:    "user_relation",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	// Prepare and send the message
	message := map[string]interface{}{
		"followerId":  followerId,
		"followeeId":  userToFollowID,
		"unfollowing": true,
		"prod":        prod,
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Println("Error marshaling Kafka message:", err)
		return
	}

	msg := kafka.Message{
		Key:   []byte(followerId),
		Value: messageJSON,
	}

	err = kafkaWriter.WriteMessages(ctx, msg)
	if err != nil {
		// Handle Kafka sending error (log or return an error)
		log.Println("Error sending message to Kafka:", err)
	}
}

func KafkaFollowRequest(ctx context.Context, followerId string, userToFollowID string, prod bool) {
	// Kafka configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{KafkaUrl},
		Topic:    "user_relation",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	// Prepare and send the message
	message := map[string]interface{}{
		"followerId":    followerId,
		"followeeId":    userToFollowID,
		"followRequest": true,
		"prod":          prod,
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Println("Error marshaling Kafka message:", err)
		return
	}

	msg := kafka.Message{
		Key:   []byte(followerId),
		Value: messageJSON,
	}

	err = kafkaWriter.WriteMessages(ctx, msg)
	if err != nil {
		// Handle Kafka sending error (log or return an error)
		log.Println("Error sending message to Kafka:", err)
	}
}

func KafkaAcceptFollowRequest(ctx context.Context, followerId string, userToFollowID string, prod bool) {
	// Kafka configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{KafkaUrl},
		Topic:    "user_relation",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	// Prepare and send the message
	message := map[string]interface{}{
		"followerId": followerId,
		"followeeId": userToFollowID,
		"accept":     true,
		"prod":       prod,
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Println("Error marshaling Kafka message:", err)
		return
	}

	msg := kafka.Message{
		Key:   []byte(followerId),
		Value: messageJSON,
	}

	err = kafkaWriter.WriteMessages(ctx, msg)
	if err != nil {
		// Handle Kafka sending error (log or return an error)
		log.Println("Error sending message to Kafka:", err)
	}
}

func KafkaDeclineFollowRequest(ctx context.Context, followerId string, userToFollowID string, prod bool) {
	// Kafka configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{KafkaUrl},
		Topic:    "user_relation",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	// Prepare and send the message
	message := map[string]interface{}{
		"followerId": followerId,
		"followeeId": userToFollowID,
		"decline":    true,
		"prod":       prod,
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Println("Error marshaling Kafka message:", err)
		return
	}

	msg := kafka.Message{
		Key:   []byte(followerId),
		Value: messageJSON,
	}

	err = kafkaWriter.WriteMessages(ctx, msg)
	if err != nil {
		// Handle Kafka sending error (log or return an error)
		log.Println("Error sending message to Kafka:", err)
	}
}

func KafkaCancelFollowRequest(ctx context.Context, followerId string, userToFollowID string, prod bool) {
	// Kafka configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{KafkaUrl},
		Topic:    "user_relation",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	// Prepare and send the message
	message := map[string]interface{}{
		"followerId":          followerId,
		"followeeId":          userToFollowID,
		"deleteFollowRequest": true,
		"prod":                prod,
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Println("Error marshaling Kafka message:", err)
		return
	}

	msg := kafka.Message{
		Key:   []byte(followerId),
		Value: messageJSON,
	}

	err = kafkaWriter.WriteMessages(ctx, msg)
	if err != nil {
		// Handle Kafka sending error (log or return an error)
		log.Println("Error sending message to Kafka:", err)
	}
}

///////////////RESTAURANTS////////////////////

func KafkaFollowRestaurant(ctx context.Context, followerId string, userToFollowID string, prod bool) {
	log.Print("KafkaUrl: ", KafkaUrl)
	// Kafka configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{KafkaUrl},
		Topic:    "user_relation",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	// Prepare and send the message
	message := map[string]interface{}{
		"followerId":          followerId,
		"followeeId":          userToFollowID,
		"followingRestaurant": true,
		"prod":                prod,
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Println("Error marshaling Kafka message:", err)
		return
	}

	msg := kafka.Message{
		Key:   []byte(followerId),
		Value: messageJSON,
	}

	err = kafkaWriter.WriteMessages(ctx, msg)
	if err != nil {
		// Handle Kafka sending error (log or return an error)
		log.Println("Error sending message to Kafka:", err)
	}
}

func KafkaUnFollowRestaurant(ctx context.Context, followerId string, userToFollowID string, prod bool) {
	// Kafka configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{KafkaUrl},
		Topic:    "user_relation",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	// Prepare and send the message
	message := map[string]interface{}{
		"followerId":            followerId,
		"followeeId":            userToFollowID,
		"unfollowingRestaurant": true,
		"prod":                  prod,
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Println("Error marshaling Kafka message:", err)
		return
	}

	msg := kafka.Message{
		Key:   []byte(followerId),
		Value: messageJSON,
	}

	err = kafkaWriter.WriteMessages(ctx, msg)
	if err != nil {
		// Handle Kafka sending error (log or return an error)
		log.Println("Error sending message to Kafka:", err)
	}
}

///////////////FEEDBACK////////////////////

func KafkaLeaveFeedbackRestaurant(ctx context.Context, userId string, restaurantId string, locationId string, reservationId string, rating uint16, feedback string, createdAt string, prod bool) {
	// Kafka configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{KafkaUrl},
		Topic:    "user_restaurant_feedback",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	// Prepare and send the message
	message := map[string]interface{}{
		"userId":        userId,
		"restaurantId":  restaurantId,
		"locationId":    locationId,
		"reservationId": reservationId,
		"rating":        rating,
		"feedback":      feedback,
		"createdAt":     createdAt,
		"prod":          prod,
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Println("Error marshaling Kafka message:", err)
		return
	}

	msg := kafka.Message{
		Key:   []byte(userId),
		Value: messageJSON,
	}

	err = kafkaWriter.WriteMessages(ctx, msg)
	if err != nil {
		// Handle Kafka sending error (log or return an error)
		log.Println("Error sending message to Kafka:", err)
	}
}
