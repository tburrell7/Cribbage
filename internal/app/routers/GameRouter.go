package routers

import (
	"fmt"
	"gotest/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)


func APIGetGames(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "JSON")
	games := service.GetGames()
	if len(games) == 0 {
		return c.String(http.StatusOK, "No Games Found")
	}
	if err := c.Bind(games); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, games)
}

func APIAddGame(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "JSON")
	game, err := service.AddGame(c.Request().Body)
	if err != nil {
		fmt.Println(err)
	}
	if err = c.Bind(game); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, game)
}

func APIRemoveGame(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "JSON")
	game, err := service.RemoveGame(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}
	if err := c.Bind(game); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, game)
}

func APIGetGame(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "JSON")
	game, err := service.GetGame(c.Param("id"))
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, game)
}

func APIUpdateScore(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "JSON")
	//game, err := service.UpdateScore(c.Param("id"), c.Request().Body)
	err := service.UpdateScore(c.Param("id"), c.Request().Body)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	return c.String(http.StatusOK, "Game Score updated")
}