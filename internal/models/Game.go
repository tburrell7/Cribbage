package models

const scoreToWin = 121

type Game struct {
	Left 		Player
	Right 		Player
	LeftScore 	int 	`bson:"leftscore"`
	RightScore	int 	`bson:"rightscore"`
}

func (game Game) isBeingPlayed() bool {
	return game.LeftScore < scoreToWin && game.RightScore < scoreToWin
}
