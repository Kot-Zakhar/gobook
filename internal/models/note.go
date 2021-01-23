package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Note struct {
	Id        primitive.ObjectID `bson:"id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Category  string             `bson:"category"`
	Subject   string             `bson:"subject"`
	Text      string             `bson:"text"`
}
