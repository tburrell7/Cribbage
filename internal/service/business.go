package service

import (
	"encoding/json"
	"gotest/internal/models"
	"io"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func GetGames() []models.Game {
	return models.GetGames()
}

func GetPlayers() []models.Player{
	return models.GetPlayers()
}

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

func RemovePlayer(body io.ReadCloser) {
	b, err := io.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}
	var p models.Player
	json.Unmarshal(b, &p)
	models.RemovePlayer(p.Id)
	return
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