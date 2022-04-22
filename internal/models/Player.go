package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Player struct {
	//ID 		primitive.ObjectID 	`bson:"_id"`
    Name 	string 				`bson:"name"`
}

type BPlayer struct {
	ID 		primitive.ObjectID 	`bson:"_id"`
    Name 	string 				`bson:"name"`
}