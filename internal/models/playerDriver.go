package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func GetPlayers() []BPlayer {
	playerCollection := client.Database("Cribbage").Collection("players")
	cursor, err := playerCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var results []BPlayer
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

func RemovePlayers() {
	collection := client.Database("Cribbage").Collection("players")
	filter := bson.D{{}}
	_, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
}