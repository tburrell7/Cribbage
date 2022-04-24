package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetGames() []Game {
	gameCollection := client.Database("Cribbage").Collection("games")
	cursor, err := gameCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var games []Game
	err = cursor.All(context.TODO(), &games)
	if err != nil {
		log.Fatal(err)
	}
	return games
}

func AddGame(left Player, right Player) (Game, error) {
	gameCollection := client.Database("Cribbage").Collection("games")
	var s Score = newScore()
	game := Game{Left: left, Right: right, Score : s}
	_, err := gameCollection.InsertOne(context.TODO(), game)
	return game, err
}

func RemoveGame(id primitive.ObjectID) {
	gameCollection := client.Database("Cribbage").Collection("games")
	filter := bson.D{{Key: "_id", Value: id}}
	_, err := gameCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
	}
}