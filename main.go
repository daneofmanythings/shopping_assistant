package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/daneofmanythings/shopping_assistant/internal/config"
	"github.com/daneofmanythings/shopping_assistant/internal/handlers"
	"github.com/daneofmanythings/shopping_assistant/internal/routes"
	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("could not load environment variables: %s", err)
	}

	port := ":" + os.Getenv("PORT")
	dbURL := os.Getenv("AUTH_DATABASE_URL")

	db, err := sql.Open("libsql", dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", dbURL, err)
		os.Exit(1)
	}
	defer db.Close()

	config := config.NewConfig()
	handlers.LinkRepoToHandlers(config)

	router := routes.NewRouter(config)

	server := http.Server{
		Addr:    port,
		Handler: router,
	}

	log.Printf("Listening on port: %s", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("error when ListenAndServe'ing: %s", err)
	}
}
