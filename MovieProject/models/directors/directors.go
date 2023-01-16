package directors

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Director struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Name string             `json:"name,omitempty"`
	Age  int                `json:"age,omitempty"`
}
