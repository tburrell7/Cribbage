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
//Player should be unique in name
func FindPlayers(filter bson.D) (Player, error) {
	playerCollection := client.Database("Cribbage").Collection("players")
	result:= playerCollection.FindOne(context.TODO(), filter)
	
	var p Player
	err := result.Decode(&p)
	return p, err
}

func FindPlayerById(id primitive.ObjectID) (Player, error) {
	var p Player
	playerCollection := client.Database("Cribbage").Collection("players")
	filter := bson.D{{Key : "_id", Value : id}}
	err := playerCollection.FindOne(context.TODO(), filter).Decode(&p)
	return p, err
}

func AddPlayer(name string) (Player, error) {
	playerCollection := client.Database("Cribbage").Collection("players")
	player := Player{Name: name}
	_, err := playerCollection.InsertOne(context.TODO(), player)
	return player, err
}

func RemovePlayer(id primitive.ObjectID) {
	playerCollection := client.Database("Cribbage").Collection("players")
	filter := bson.D{{Key: "_id", Value: id}}
	_, err := playerCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
}