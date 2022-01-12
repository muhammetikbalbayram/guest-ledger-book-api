package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Guest struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Mail    string             `json:"mail"`
	Message string             `json:"message"`
}
