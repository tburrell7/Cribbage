package app

import (
	"fmt"
	"gotest/internal/app/routers"
	"log"
	"net/http"
)

func Run() {
	http.HandleFunc("/games", routers.GameRouter)
	http.HandleFunc("/players", routers.PlayerRouter)
	
	fmt.Printf("Starting server at port 3000\n")
    if err := http.ListenAndServe(":3000", nil); err != nil {
        log.Fatal(err)
    }
	
	

	// fmt.Println("Let's Play!")
	// leftTest := models.NewPlayer("Thomas")
	// rightTest := models.NewPlayer("Josh")
	// models.PrintGames(client)
	// id := models.NewGame(client, leftTest, rightTest)
	// models.PrintGames(client)
	// models.UpdateScore(client, id, 60, 20)
	// models.PrintGames(client)
	// models.UpdateScore(client, id, 121, 110)
	// models.PrintGames(client)
	// models.RemoveGame(client, id)
	// models.PrintGames(client)
}