package service

import (
	"fmt"
	"gotest/internal/models"
)

func AddGame() {
	Thomas:= models.NewPlayer("Thomas")
	Josh:= models.NewPlayer("Josh")
	models.NewGame(Thomas, Josh)
}

func ClearData() {
	models.RemoveGames()
	models.RemovePlayers()
}

func PrintGames() string {
	games := models.GetGames()
	var s = ""
	for i := 0; i < len(games); i++ {
		s += fmt.Sprintf("Left %s vs Player %s\n", games[i].Left.Name, games[i].Right.Name)
	}
	return s
}

func GetPlayers() []models.Player{
	return models.GetPlayers()
}