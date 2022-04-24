package models

import "go.mongodb.org/mongo-driver/bson/primitive"

const scoreToWin = 121

type Game struct {
	Id			primitive.ObjectID	`bson:"_id,omitempty"`
	Left 		Player				`bson:"left"`
	Right 		Player				`bson:"right"`
	Score		Score
}

func (game Game) isBeingPlayed() bool {
	return game.Score.Left < scoreToWin && game.Score.Right < scoreToWin
}
