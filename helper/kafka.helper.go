package helper

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

func KafkaFollow(ctx context.Context, userId string, userToFollowID string) {
	// Kafka configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "user_relation",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	// Prepare and send the message
	message := map[string]interface{}{
		"userId":      userId,
		"follow_user": userToFollowID,
		"following":   true,
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

func KafkaUnFollow(ctx context.Context, userId string, userToFollowID string) {
	// Kafka configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "user_relation",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	// Prepare and send the message
	message := map[string]interface{}{
		"userId":      userId,
		"follow_user": userToFollowID,
		"unfollowing": true,
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

func KafkaFollowRequest(ctx context.Context, userId string, userToFollowID string) {
	// Kafka configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "user_relation",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	// Prepare and send the message
	message := map[string]interface{}{
		"userId":        userId,
		"follow_user":   userToFollowID,
		"followRequest": true,
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

func KafkaAcceptFollowRequest(ctx context.Context, userId string, userToFollowID string) {
	// Kafka configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "user_relation",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	// Prepare and send the message
	message := map[string]interface{}{
		"userId":      userId,
		"follow_user": userToFollowID,
		"accept":      true,
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

func KafkaDeclineFollowRequest(ctx context.Context, userId string, userToFollowID string) {
	// Kafka configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "user_relation",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	// Prepare and send the message
	message := map[string]interface{}{
		"userId":      userId,
		"follow_user": userToFollowID,
		"decline":     true,
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

///////////////RESTAURANTS////////////////////

func KafkaFollowRestaurant(ctx context.Context, userId string, userToFollowID string) {
	// Kafka configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "user_relation",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	// Prepare and send the message
	message := map[string]interface{}{
		"userId":              userId,
		"follow_user":         userToFollowID,
		"followingRestaurant": true,
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

func KafkaUnFollowRestaurant(ctx context.Context, userId string, userToFollowID string) {
	// Kafka configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "user_relation",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	// Prepare and send the message
	message := map[string]interface{}{
		"userId":                userId,
		"follow_user":           userToFollowID,
		"unfollowingRestaurant": true,
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
