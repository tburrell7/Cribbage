package routers

import (
	"fmt"
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

    fmt.Fprintf(w, "Players!")
}