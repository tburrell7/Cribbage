package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func newScore() Score {
	var s Score = Score{Left : 0, Right : 0}
	return s
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