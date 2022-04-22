package routers

import (
	"fmt"
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
		fmt.Fprintf(w, "Games!")
		listOfGames := service.PrintGames()
		fmt.Fprintf(w, listOfGames)

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
