package person

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string             `json:"title,omitempty"`
	Name  string             `json:"name,omitempty"`
	Age   int                `json:"age,omitempty"`
	DoB   string             `json:"dob,omitempty"`
}
