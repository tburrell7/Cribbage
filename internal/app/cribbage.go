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
}