package actor

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Actor struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty"`
	Age  int                `json:"age,omitempty"`
}
