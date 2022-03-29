package app

import (
	"fmt"
	"gotest/internal/models"
)
func newGame(left models.Player, right models.Player) models.Game {
	return models.Game{Left: left, Right: right, Ongoing: true, LeftScore: 0, RightScore: 0}
}

func Run() {
	fmt.Println("Let's Play!")
	leftTest := models.Player{Name: "Thomas"}
	rightTest := models.Player{Name: "Josh"}
	gameTest := newGame(leftTest, rightTest)
	gameTest.UpdateScore(60, 20)
	gameTest.PrintStats()
	gameTest.UpdateScore(121, 110)
	gameTest.PrintStats()
}