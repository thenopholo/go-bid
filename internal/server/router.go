package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/thenopholo/go-bid/internal/handler"
)

func Router(r *chi.Mux, h *handler.Handler) {
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/healthcheck", handler.HealthCheck)

			r.Route("/users", func(r chi.Router) {
				r.Post("/signup", h.UserSignup)
				r.Post("/login", h.UserLogin)
				r.Post("/logout", h.UserLogin)
			})
		})
	})
}