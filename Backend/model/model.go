package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Task   string             `bson:"task" json:"task"`
	Status bool               `bson:"status" json:"status"`
}
