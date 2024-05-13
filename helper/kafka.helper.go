package helper

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

func KafkaFollow(ctx context.Context, followerId string, userToFollowID string) {
	// Kafka configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "user_relation",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	// Prepare and send the message
	message := map[string]interface{}{
		"followerId": followerId,
		"followeeId": userToFollowID,
		"following":  true,
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

func KafkaUnFollow(ctx context.Context, followerId string, userToFollowID string) {
	// Kafka configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "user_relation",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	// Prepare and send the message
	message := map[string]interface{}{
		"followerId":  followerId,
		"followeeId":  userToFollowID,
		"unfollowing": true,
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

func KafkaFollowRequest(ctx context.Context, followerId string, userToFollowID string) {
	// Kafka configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "user_relation",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	// Prepare and send the message
	message := map[string]interface{}{
		"followerId":    followerId,
		"followeeId":    userToFollowID,
		"followRequest": true,
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

func KafkaAcceptFollowRequest(ctx context.Context, followerId string, userToFollowID string) {
	// Kafka configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "user_relation",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	// Prepare and send the message
	message := map[string]interface{}{
		"followerId": followerId,
		"followeeId": userToFollowID,
		"accept":     true,
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

func KafkaDeclineFollowRequest(ctx context.Context, followerId string, userToFollowID string) {
	// Kafka configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "user_relation",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	// Prepare and send the message
	message := map[string]interface{}{
		"followerId": followerId,
		"followeeId": userToFollowID,
		"decline":    true,
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

func KafkaCancelFollowRequest(ctx context.Context, followerId string, userToFollowID string) {
	// Kafka configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "user_relation",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	// Prepare and send the message
	message := map[string]interface{}{
		"followerId":          followerId,
		"followeeId":          userToFollowID,
		"deleteFollowRequest": true,
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

func KafkaFollowRestaurant(ctx context.Context, followerId string, userToFollowID string) {
	// Kafka configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "user_relation",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	// Prepare and send the message
	message := map[string]interface{}{
		"followerId":          followerId,
		"followeeId":          userToFollowID,
		"followingRestaurant": true,
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

func KafkaUnFollowRestaurant(ctx context.Context, followerId string, userToFollowID string) {
	// Kafka configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "user_relation",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	// Prepare and send the message
	message := map[string]interface{}{
		"followerId":            followerId,
		"followeeId":            userToFollowID,
		"unfollowingRestaurant": true,
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
