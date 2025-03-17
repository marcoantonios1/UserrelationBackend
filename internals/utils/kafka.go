package utils

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

func writeMessageToKafka(ctx context.Context, topic string, key string, message interface{}) error {
	KafkaUrl := os.Getenv("KAFKA_URL")
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{KafkaUrl},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer kafkaWriter.Close()

	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Println("Error marshaling Kafka message:", err)
		return err
	}

	msg := kafka.Message{
		Key:   []byte(key),
		Value: messageJSON,
	}

	if err := kafkaWriter.WriteMessages(ctx, msg); err != nil {
		log.Println("Error sending message to Kafka:", err)
		return err
	}

	return nil
}

func KafkaFollow(ctx context.Context, followerId string, userToFollowID string, prod bool) {
	message := map[string]interface{}{
		"followerId": followerId,
		"followeeId": userToFollowID,
		"following":  true,
		"prod":       prod,
	}

	_ = writeMessageToKafka(ctx, "user_relation", followerId, message)
}

func KafkaUnFollow(ctx context.Context, followerId string, userToFollowID string, prod bool) {
	message := map[string]interface{}{
		"followerId":  followerId,
		"followeeId":  userToFollowID,
		"unfollowing": true,
		"prod":        prod,
	}

	_ = writeMessageToKafka(ctx, "user_relation", followerId, message)
}

func KafkaFollowRequest(ctx context.Context, followerId string, userToFollowID string, prod bool) {
	message := map[string]interface{}{
		"followerId":    followerId,
		"followeeId":    userToFollowID,
		"followRequest": true,
		"prod":          prod,
	}

	_ = writeMessageToKafka(ctx, "user_relation", followerId, message)
}

func KafkaAcceptFollowRequest(ctx context.Context, followerId string, userToFollowID string, prod bool) {
	message := map[string]interface{}{
		"followerId": followerId,
		"followeeId": userToFollowID,
		"accept":     true,
		"prod":       prod,
	}

	_ = writeMessageToKafka(ctx, "user_relation", followerId, message)
}

func KafkaDeclineFollowRequest(ctx context.Context, followerId string, userToFollowID string, prod bool) {
	message := map[string]interface{}{
		"followerId": followerId,
		"followeeId": userToFollowID,
		"decline":    true,
		"prod":       prod,
	}

	_ = writeMessageToKafka(ctx, "user_relation", followerId, message)
}

func KafkaCancelFollowRequest(ctx context.Context, followerId string, userToFollowID string, prod bool) {
	message := map[string]interface{}{
		"followerId":          followerId,
		"followeeId":          userToFollowID,
		"deleteFollowRequest": true,
		"prod":                prod,
	}

	_ = writeMessageToKafka(ctx, "user_relation", followerId, message)
}

func KafkaFollowLog(ctx context.Context, userId string, targetUserId string, eventType string, prod bool) {
	message := map[string]interface{}{
		"userId":       userId,
		"targetUserId": targetUserId,
		"eventType":    eventType,
		"userFollow":   true,
		"prod":         prod,
	}

	_ = writeMessageToKafka(ctx, "user_action_log", userId, message)
}

///////////////RESTAURANTS////////////////////

func KafkaFollowRestaurant(ctx context.Context, followerId string, userToFollowID string, prod bool) {
	message := map[string]interface{}{
		"followerId":          followerId,
		"followeeId":          userToFollowID,
		"followingRestaurant": true,
		"prod":                prod,
	}

	_ = writeMessageToKafka(ctx, "user_relation", followerId, message)
}

func KafkaUnFollowRestaurant(ctx context.Context, followerId string, userToFollowID string, prod bool) {
	message := map[string]interface{}{
		"followerId":            followerId,
		"followeeId":            userToFollowID,
		"unfollowingRestaurant": true,
		"prod":                  prod,
	}

	_ = writeMessageToKafka(ctx, "user_relation", followerId, message)
}

func KafkaRestaurantFollowLog(ctx context.Context, userId string, restaurantId string, eventType string, prod bool) {
	message := map[string]interface{}{
		"userId":           userId,
		"restaurantId":     restaurantId,
		"eventType":        eventType,
		"restaurantFollow": true,
		"prod":             prod,
	}

	_ = writeMessageToKafka(ctx, "user_action_log", userId, message)
}

///////////////FEEDBACK////////////////////

func KafkaLeaveFeedbackRestaurant(ctx context.Context, userId string, restaurantId string, locationId string, reservationId string, rating uint16, feedback string, createdAt string, prod bool) {
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

	_ = writeMessageToKafka(ctx, "user_restaurant_feedback", userId, message)
}
