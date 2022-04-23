package routers

import (
	"encoding/json"
	"fmt"
	"log"

	//"gotest/internal/app"
	"gotest/internal/service"
	"net/http"
)

func GameRouter(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/games" {
		//http.Error(w, "404 not found.", http.StatusNotFound)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		games := service.GetGames()
		if len(games) == 0 {
			w.WriteHeader(http.StatusNoContent)
			fmt.Fprint(w, "No Games Avaliable")
		}
		resp, err := json.Marshal(games)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(resp)
		return
	} else if r.Method == "PUT" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	} else if r.Method == "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	} else if r.Method == "DELETE" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
}
