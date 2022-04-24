package models

import "go.mongodb.org/mongo-driver/bson/primitive"

const scoreToWin = 121

type Game struct {
	ID			primitive.ObjectID	`bson:"_id,omitempty"`
	Left 		Player				`bson:"left"`
	Right 		Player				`bson:"right"`
	LeftScore 	int 				`bson:"leftscore"`
	RightScore	int 				`bson:"rightscore"`
}

func (game Game) isBeingPlayed() bool {
	return game.LeftScore < scoreToWin && game.RightScore < scoreToWin
}
