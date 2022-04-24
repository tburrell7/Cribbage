package service

import (
	"encoding/json"
	"fmt"
	"gotest/internal/models"
	"io"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func GetGames() []models.Game {
	return models.GetGames()
}

func AddGame(body io.ReadCloser) (models.Game, error) {
	var err error
	b, err := io.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}
	var g models.Game
	json.Unmarshal(b, &g)
	pL := g.Left
	pR := g.Right
	if b, e := playerExists(pL.Name); b {
		pL, e = getPlayerByName(g.Left.Name)
		err = e
	} else {
		err = e
	}
		
	if b, e := playerExists(pL.Name); b {
		pL, e = getPlayerByName(g.Left.Name)
		err = e
	} else {
		err = e
	}
	if err == nil {
		g, err = models.AddGame(pL, pR)
	}
	return g, err
}

func GetPlayers() []models.Player{
	return models.GetPlayers()
}

func getPlayerByName(name string) (models.Player, error) {
	player, err := models.FindPlayers(bson.D{{Key: "name", Value: name}})
	return player, err
}

func AddPlayer(body io.ReadCloser) (models.Player, error) {
	b, err := io.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}
	var p models.Player
	json.Unmarshal(b, &p)
	if b, _ := playerExists(p.Name); !b {
		_, err = models.AddPlayer(p.Name)
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

func playerExists(name string) (bool, error) {
	b := true
	_, err := models.FindPlayers(bson.D{{Key: "name", Value: name}})
	if err != nil {
		b = false
	}
	fmt.Println(name, "is a player")
	return b, err
}