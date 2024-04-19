package handlers

import (
	"fmt"
	"net/http"

	"github.com/daneofmanythings/shopping_assistant/internal/config"
)

var R *config.Repo

func LinkRepoToHandlers(c *config.Config) {
	R = &c.Repo
}

func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func GetPathParameters(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("Path param recieved: " + id))
}
