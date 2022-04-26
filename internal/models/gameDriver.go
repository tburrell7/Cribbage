package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func newGame() Game {
// 	var p1, p2 Player
// 	var s Score
// 	p1.Name = ""
// 	p2.Name = ""
// 	s.Left = 0
// 	s.Right = 0
// 	var g Game = Game{Left : p1, Right : p2, Score: s}
// 	return g
// }

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

func GetGameById(id primitive.ObjectID) (Game, error) {
	gameCollection := client.Database("Cribbage").Collection("games")
	var g Game
	filter := bson.D{{Key : "_id", Value : id}}
	err := gameCollection.FindOne(context.TODO(), filter).Decode(&g)
	fmt.Print("models.GetGameById() ", g.Id, id)
	if err != nil {
		fmt.Println("models.GetGameById() ", g.Id, id)
		fmt.Println(err)
	}
	return g, err
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