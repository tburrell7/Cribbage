package models

import "fmt"

const ScoreToWin = 121
type Game struct {
    Left Player
	Right Player
	Ongoing bool
	LeftScore int
	RightScore int
}

func (game *Game) UpdateScore(leftScore int, rightScore int) {
	game.LeftScore = leftScore
	game.RightScore = rightScore
	if leftScore >= ScoreToWin || rightScore >= ScoreToWin {
		game.Ongoing = false
	}
}

func (game Game) PrintStats() {
	if game.Ongoing {
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