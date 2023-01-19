package directors

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Director struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty" validate:"required"`
	Age  int                `json:"age,omitempty" validate:"required,numeric"`
	DoB  string             `json:"dob,omitempty" validate:"required"`
}
