package service

import (
	"encoding/json"
	"fmt"
	"gotest/internal/models"
	"io"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func AddPlayer(body io.ReadCloser) (models.Player, error) {
	b, err := io.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}
	var p models.Player
	json.Unmarshal(b, &p)
	if !playerExists(p.Name) {
		err = models.AddPlayer(p.Name)
	}
	return p, err
}

func playerExists(name string) bool {
	players := models.FindPlayers(bson.D{{Key: "name", Value: name}})
	for i := 0; i < len(players); i++ {
		if players[i].Name == name {
			return true
		}
	}
	return false
}

func PrintGames() string {
	games := models.GetGames()
	var s = ""
	for i := 0; i < len(games); i++ {
		s += fmt.Sprintf("Left %s vs Player %s\n", games[i].Left.Name, games[i].Right.Name)
	}
	return s
}

func GetPlayers() []models.BPlayer{
	return models.GetPlayers()
}