package model

import (
	"go.mongodb.org/mongo-driver/mongo/bson/primitive"

)

type Movie struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie string             `json:"movie,omitempty" bson:"movie,omitempty"`
	Watched bool             `json:"watched,omitempty" bson:"watched,omitempty"`
}
