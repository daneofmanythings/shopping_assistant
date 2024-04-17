package routes

import (
	"github.com/daneofmanythings/shopping_assistant/internal/config"
	"github.com/daneofmanythings/shopping_assistant/internal/handlers"
	v1 "github.com/daneofmanythings/shopping_assistant/internal/routes/v1"
	chi "github.com/go-chi/chi/v5"
)

func NewRouter(c *config.Config) *chi.Mux {
	m := chi.NewMux()

	m.Get("/", handlers.GetRoot)

	m.Mount("/v1", v1.NewV1Router(c))

	return m
}
