package routes

import (
	"github.com/go-chi/chi/v5"
	handler "main.go/app/handlers"
)

func ApiRouters(r chi.Router) {
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/health", handler.ServerHealth)
		})
	})
}