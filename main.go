package main

import (
	"gotest/internal/app"
	"gotest/internal/models"
)

func main() {
	models.ConnectToMongo()
	defer models.DisconnectFromMongo()
	app.Run()
}