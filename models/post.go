package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Post model info
// @Description Post information
type Post struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty" swaggertype:"primitive,string"`
	Title   string             `bson:"title" json:"title"`
	Content string             `bson:"content" json:"content"`
	Author  string             `bson:"author" json:"author"`
}
