package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
// Todocollection ...
	Todocollection = "Todos"
)

// Todo ...
type Todo struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Title string `json:"title"`
	Description string `json:"description"`
}