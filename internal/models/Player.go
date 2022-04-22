package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Player struct {
	ID 		primitive.ObjectID 	`bson:"_id"`
    Name 	string 				`bson:"name"`
}

func NewPlayer(Name string) Player {
	collection := client.Database("Cribbage").Collection("players")
	player := Player{Name: Name}
	_, err := collection.InsertOne(context.TODO(), player)
	if err != nil {
		log.Fatal(err)
	}
	return player
}

func RemovePlayer(id interface{}) {
	collection := client.Database("Cribbage").Collection("players")
	filter := bson.D{{"_id", id}}
	deleteRes, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted", deleteRes.DeletedCount, "player with ID =", id)
}