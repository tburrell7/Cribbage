package app

import (
	"fmt"
	"gotest/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func Run(client *mongo.Client) {
	fmt.Println("Let's Play!")
	leftTest := models.NewPlayer("Thomas")
	rightTest := models.NewPlayer("Josh")
	models.PrintGames(client)
	id := models.NewGame(client, leftTest, rightTest)
	models.PrintGames(client)
	models.RemoveGame(client, id)
	models.PrintGames(client)
	// gameTest.UpdateScore(60, 20)
	// gameTest.PrintStats()
	// gameTest.UpdateScore(121, 110)
	// gameTest.PrintStats()
}