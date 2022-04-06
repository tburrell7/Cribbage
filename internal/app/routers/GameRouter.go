package routers

import (
	"fmt"
	"net/http"
)

func GameRouter(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/games" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }

    fmt.Fprintf(w, "Games!")
}