package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Challenge struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Name   string             `json:"name" bson:"name"`
	Tasks  []Task             `json:"tasks" bson:"tasks"`
	IsDone bool               `json:"isDone" bson:"isDone"`
}
