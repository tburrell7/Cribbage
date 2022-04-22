package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func ConnectToMongo() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	mongoClient, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = mongoClient.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	client = mongoClient
	// collection := client.Database("profile").Collection("people")

	// tom := models.Person{Name:"Tom", Age:21}
	// insertResult, err := collection.InsertOne(context.TODO(), tom)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// filter := bson.D{{"first_name", "Thomas"}}
	// update := bson.D{
	// 	{"$inc", bson.D{
	// 		{"age", 1},
	// 	}},
	// }

	// updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.UpsertedCount)
}

func DisconnectFromMongo() {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
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
	filter := bson.D{{"_id", id}}
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

func RemovePlayers() {
	collection := client.Database("Cribbage").Collection("players")
	filter := bson.D{{}}
	deleteRes, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted", deleteRes.DeletedCount, "players")
}

func UpdateScore(id interface{}, leftScore int, rightScore int) interface{} {
	collection := client.Database("Cribbage").Collection("games")
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"leftscore", leftScore}, {"rightscore", rightScore}}}}
	updateRes, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return updateRes.UpsertedID
}

func GetGames() []Game {
	collection := client.Database("Cribbage").Collection("games")
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var games []Game
	if err = cursor.All(context.TODO(), &games); err != nil {
		log.Fatal(err)
	}
	return games

	// if len(results) == 0 {
	// 	fmt.Fprintf(w, "No Games Exist")
	// }
	// for _, result := range results {
	// 	fmt.Fprintf(w, GameToString(result))
	// }
}

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
