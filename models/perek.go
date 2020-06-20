package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Perek struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Gender  int
	Tattoos []Tattoo
}
