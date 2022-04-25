package service

import (
	"encoding/json"
	"gotest/internal/models"
	"io"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	if b, e := playerExists(pR.Name); b {
		pR, e = getPlayerByName(g.Right.Name)
		err = e
	} else {
		err = e
	}
	if err == nil {
		g, err = models.AddGame(pL, pR)
	}
	return g, err
}

func RemoveGame(body io.ReadCloser) models.Game {
	b, err := io.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}
	var g models.Game
	json.Unmarshal(b, &g)
	models.RemoveGame(g.Id)
	return g
}

func GetPlayers() []models.Player {
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

func RemovePlayer(body io.ReadCloser) models.Player {
	b, err := io.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}
	var p models.Player
	json.Unmarshal(b, &p)
	models.RemovePlayer(p.Id)
	return p
}

func GetPlayer(id primitive.ObjectID) (models.Player, error) {
	return models.FindPlayerById(id)
}

func playerExists(name string) (bool, error) {
	b := true
	_, err := models.FindPlayers(bson.D{{Key: "name", Value: name}})
	if err != nil {
		b = false
	}
	return b, err
}
