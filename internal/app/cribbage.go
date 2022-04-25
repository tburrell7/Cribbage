package app

import (
	"gotest/internal/app/routers"
	//"net/http"

	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()
	e.GET("/games", routers.APIGetGames)
	e.POST("/games", routers.APIAddGame)
	e.DELETE("/games", routers.APIRemoveGame)
	e.GET("/players", routers.APIGetPlayers)
	e.POST("/players", routers.APIAddPlayer)
	e.DELETE("/players", routers.APIRemovePlayer)

	e.Logger.Fatal(e.Start(":3000"))
}