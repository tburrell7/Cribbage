package app

import (
	"gotest/internal/app/routers"
	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()
	e.GET("/games", routers.APIGetGames)
	e.POST("/games", routers.APIAddGame)

	e.GET("/games/:id", routers.APIGetGame)
	e.PATCH("/games/:id", routers.APIUpdateScore)
	e.DELETE("/games/:id", routers.APIRemoveGame)

	e.GET("/players", routers.APIGetPlayers)
	e.POST("/players", routers.APIAddPlayer)
	e.DELETE("/players", routers.APIRemovePlayer)

	e.GET("/players/:id", routers.APIGetPlayer)

	e.Logger.Fatal(e.Start(":3000"))
}