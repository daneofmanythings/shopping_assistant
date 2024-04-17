package v1

import (
	"github.com/daneofmanythings/shopping_assistant/internal/config"
	"github.com/daneofmanythings/shopping_assistant/internal/handlers"
	chi "github.com/go-chi/chi/v5"
)

func NewV1Router(c *config.Config) *chi.Mux {
	m := chi.NewMux()

	m.Get("/param/{id}", handlers.GetPathParameters)

	return m
}
