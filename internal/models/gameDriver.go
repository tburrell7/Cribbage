package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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

func NewGame(left Player, right Player) interface{} {
	collection := client.Database("Cribbage").Collection("games")
	game := Game{Left: left, Right: right, LeftScore: 0, RightScore: 0}
	insertResult, err := collection.InsertOne(context.TODO(), game)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult.InsertedID
}

func RemoveGame(id interface{}) {
	collection := client.Database("Cribbage").Collection("games")
	filter := bson.D{{Key: "_id", Value: id}}
	deleteRes, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted", deleteRes.DeletedCount, "games with ID =", id)
}

func RemoveGames() {
	collection := client.Database("Cribbage").Collection("games")
	filter := bson.D{{}}
	deleteRes, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted", deleteRes.DeletedCount, "games")
}

func UpdateScore(id interface{}, leftScore int, rightScore int) interface{} {
	collection := client.Database("Cribbage").Collection("games")
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "leftscore", Value: leftScore}, {Key: "rightscore", Value: rightScore}}}}
	updateRes, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return updateRes.UpsertedID
}

func PrintStats(game Game) {
	if game.isBeingPlayed() {
		fmt.Println(game.Left.Name, "vs", game.Right.Name, "\b! The score is", game.LeftScore, "to", game.RightScore, "\b!")
	} else {
		leftWin := game.LeftScore > game.RightScore
		if leftWin {
			fmt.Println(game.Left.Name, "beats", game.Right.Name, "\b! With a score of", game.LeftScore, "to", game.RightScore, "\b!")
		} else {
			fmt.Println(game.Right.Name, "beats", game.Left.Name, "\b! With a score of", game.RightScore, "to", game.LeftScore, "\b!")
		}
	}
}
