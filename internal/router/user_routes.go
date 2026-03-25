package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/rana-touseef11/go-chi-postgresql/internal/handler"
	"github.com/rana-touseef11/go-chi-postgresql/internal/middleware"
)

func RegisterUserRoutes(r chi.Router, handler *handler.UserHandler) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", handler.Login)
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", handler.GetAll)
		r.Get("/{id}", handler.GetById)

		r.Group(func(r chi.Router) {
			r.Use(middleware.JWTAuth)

			r.Post("/", handler.Create)
			r.Put("/{id}", handler.Update)
			r.Delete("/{id}", handler.Delete)
		})
	})
}
