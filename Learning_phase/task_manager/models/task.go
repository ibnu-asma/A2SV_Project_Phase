package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title" binding:"required"`
	Description string             `json:"description" bson:"description" binding:"required"`
	DueDate     time.Time          `json:"due_date" bson:"due_date" binding:"required"`
	Status      string             `json:"status" bson:"status" binding:"required"`
}
