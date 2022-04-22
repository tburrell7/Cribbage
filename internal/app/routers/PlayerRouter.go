package routers

import (
	"encoding/json"
	"fmt"
	"gotest/internal/service"
	"log"
	"net/http"
)

func PlayerRouter(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/players" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		players := service.GetPlayers()
		if len(players) == 0 {
			w.WriteHeader(http.StatusNoContent)
			fmt.Fprint(w, "No Players Avaliable")
		}
		resp, err := json.Marshal(players)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(resp)
		return
	}
}
