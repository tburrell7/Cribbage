package models

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func newScore() Score {
	var s Score = Score{LeftScore : 0, RightScore : 0}
	return s
}

func UpdateScore(id primitive.ObjectID, s Score) error {
	gameCollection := client.Database("Cribbage").Collection("games")
	filter := bson.D{{Key: "_id", Value: id}}
	update := 	bson.D{{Key: "$set", Value:
					bson.D{{Key: "score", Value:
						bson.D{{Key: "leftscore", Value: s.LeftScore}, {Key: "rightscore", Value: s.RightScore}}}}}}
	res, err := gameCollection.UpdateOne(context.TODO(), filter, update)
	fmt.Println("ID =", res.UpsertedID)
	return err
}