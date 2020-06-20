package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID                     primitive.ObjectID     `json:"id" bson:"_id"`
	Description            string                 `json:"description" bson:"description"`
	CompletionRequirements map[string]interface{} `json:"completionRequirements" bson:"completionRequirements"`
	IsDone                 bool                   `json:"isDone" bson:"isDone"`
}
