package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID               primitive.ObjectID `json:"id" bson:"_id"`
	PhoneNumber      string             `json:"phoneNumber" bson:"phoneNumber"`
	ReferalCardID    string             `json:"referalCardID" bson:"referalCardID"`
	Perek            Perek              `json:"perek" bson:"perek"`
	CurrentChallenge Challenge          `json:"currentChallenge" bson:"currentChallenge"`
	Name             string             `json:"name" bson:"name"`
	Age              int                `json:"age" bson:"age"`
	Gender           bool               `json:"gender" bson:"gender"`
}
