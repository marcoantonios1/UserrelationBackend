package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	UserName       *string            `json:"username" bson:"username" validate:"required,min=4,max=30"`
	Name           *string            `json:"name" bson:"name"`
	Image          *string            `json:"image" bson:"image"`
	Password       *string            `json:"password" bson:"password" validate:"required,min=8"`
	Email          *string            `json:"email" bson:"email" validate:"email,required"`
	Phone          *string            `json:"phone" bson:"phone" validate:"required"`
	Gender         *string            `json:"gender" bson:"gender"`
	Biography      *string            `json:"bio" bson:"bio"`
	Followers      uint32             `json:"followers" bson:"followers"`
	Following      uint32             `json:"following" bson:"following"`
	Follow_Request uint32             `json:"follow_requests" bson:"follow_requests"`
	Status         bool               `json:"status" bson:"status"`
	Private        bool               `json:"private" bson:"private"`
	Token          *string            `json:"access_token" bson:"-"`
	Refresh_Token  *string            `json:"refresh_token" bson:"-"`
	Email_Verified bool               `json:"email_verified" bson:"email_verified"`
	Birthday       time.Time          `json:"birthday" bson:"birthday"`
	Created_At     time.Time          `json:"created_at" bson:"created_at"`
	Updated_At     time.Time          `json:"updated_at" bson:"updated_at"`
}

type Neo4jUser struct {
	ID        string `json:"_id"`
	UserName  string `json:"username" validate:"required,min=4,max=30"`
	Name      string `json:"name"`
	Image     string `json:"image"`
	Biography string `json:"bio"`
	Followers uint32 `json:"followers"`
	Following uint32 `json:"following"`
	Private   bool   `json:"private"`
	Relation  string `json:"relation"`
}

type Restaurant struct {
	Restaurant_ID   string `json:"_id" bson:"_id"`
	Restaurant_Name string `json:"restaurant_name" bson:"restaurant_name" validate:"required,min=2,max=30"`
	Description     string `json:"description" bson:"description" validate:"required,min=2,max=30"`
	Image           string `json:"image" bson:"image"`
}

type Feedback struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	User_ID        primitive.ObjectID `json:"user_id" bson:"user_id"`
	Restaurant_ID  primitive.ObjectID `json:"restaurant_id" bson:"restaurant_id"`
	Location_ID    primitive.ObjectID `json:"location_id" bson:"location_id"`
	Reservation_ID primitive.ObjectID `json:"reservation_id" bson:"reservation_id"`
	Feedback       string             `json:"feedback,omitempty" bson:"feedback,omitempty"`
	Rating         uint32             `json:"rating" bson:"rating"`
	Created_At     time.Time          `json:"created_at" bson:"created_at"`
}
