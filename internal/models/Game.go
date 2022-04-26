package models

import "go.mongodb.org/mongo-driver/bson/primitive"

const scoreToWin = 121

type Game struct {
	Id			primitive.ObjectID	`bson:"_id,omitempty"`
	Left 		Player				`bson:"left"`
	Right 		Player				`bson:"right"`
	Score		Score				`bson:"score"`
}

func (game Game) isBeingPlayed() bool {
	return game.Score.LeftScore < scoreToWin && game.Score.RightScore < scoreToWin
}
