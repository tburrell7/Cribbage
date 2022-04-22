package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Player struct {
	Id 		primitive.ObjectID 	`bson:"_id,omitempty"`
    Name 	string 				`bson:"name"`
}