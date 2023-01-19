package movies

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID        primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie     string               `json:"movie,omitempty" validate:"required"`
	Watched   bool                 `json:"watched,omitempty"`
	Actors_id []primitive.ObjectID `json:"actors_id,omitempty" validate:"required"`
	Director  primitive.ObjectID   `json:"director,omitempty" validate:"required"`
	Rating    int                  `json:"rating,omitempty" validate:"required,numeric,min=0,max=10"`
}
