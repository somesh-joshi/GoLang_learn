package person

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string             `json:"title,omitempty" validate:"required d"`
	Name  string             `json:"name,omitempty" validate:"required"`
	Age   int                `json:"age,omitempty" validate:"required,numeric"`
	DoB   string             `json:"dob,omitempty" validate:"required"`
}
