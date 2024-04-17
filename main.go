package main

import (
	"log"
	"net/http"
	"os"

	"github.com/daneofmanythings/shopping_assistant/internal/config"
	"github.com/daneofmanythings/shopping_assistant/internal/handlers"
	"github.com/daneofmanythings/shopping_assistant/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("could not load environment variables: %s", err)
	}

	port := ":" + os.Getenv("PORT")

	config := &config.Config{}
	handlers.LinkConfigToHandlers(config)

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
