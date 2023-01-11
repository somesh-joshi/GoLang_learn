package movies

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID       primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie    string               `json:"movie,omitempty"`
	Watched  bool                 `json:"watched,omitempty"`
	Actors   []primitive.ObjectID `json:"actors,omitempty"`
	Director primitive.ObjectID   `json:"director,omitempty"`
}