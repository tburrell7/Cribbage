package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPlayers() []Player {
	playerCollection := client.Database("Cribbage").Collection("players")
	cursor, err := playerCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var results []Player
	cursor.All(context.TODO(), &results)
	return results
}

func FindPlayers(filter bson.D) []Player {
	playerCollection := client.Database("Cribbage").Collection("players")
	cursor, err := playerCollection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	var results []Player
	cursor.All(context.TODO(), &results)
	return results
}

func AddPlayer(name string) error {
	playerCollection := client.Database("Cribbage").Collection("players")
	player := Player{Name: name}
	_, err := playerCollection.InsertOne(context.TODO(), player)
	return err
}

func RemovePlayer(id primitive.ObjectID) {
	playerCollection := client.Database("Cribbage").Collection("players")
	filter := bson.D{{Key: "_id", Value: id}}
	_, err := playerCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
}