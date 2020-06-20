package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tattoo struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Name string
	Bone int
}
