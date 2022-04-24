package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Score struct {
	Id 		primitive.ObjectID 	`bson:"_id,omitempty"`
    Left 	int 				`bson:"left"`
	Right 	int 				`bson:"right"`
}