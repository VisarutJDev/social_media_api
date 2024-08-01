package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User model info
// @Description User information
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty" swaggertype:"primitive,string"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
}

// LoginInput model info
// @Description LoginInput information
type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse model info
// @Description AuthResponse information
type AuthResponse struct {
	Token string `json:"token"` // Response message
}
