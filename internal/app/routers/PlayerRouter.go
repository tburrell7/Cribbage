package routers

import (
	"fmt"
	"gotest/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func APIGetPlayers(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "JSON")
	players := service.GetPlayers()
	if len(players) == 0 {
		return c.String(http.StatusOK, "No Players Found")
	}
	if err := c.Bind(players); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, players)
}

func APIAddPlayer(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "JSON")
	player, err := service.AddPlayer(c.Request().Body)
	if err != nil {
		fmt.Println(err)
	}
	if err = c.Bind(player); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, player)
}

func APIRemovePlayer(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "JSON")
	player := service.RemovePlayer(c.Request().Body)
	if err := c.Bind(player); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, player)
}