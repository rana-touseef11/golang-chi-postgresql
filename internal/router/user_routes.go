package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/rana-touseef11/go-chi-postgresql/internal/dto"
	"github.com/rana-touseef11/go-chi-postgresql/internal/handler"
	"github.com/rana-touseef11/go-chi-postgresql/internal/middleware"
	"github.com/rana-touseef11/go-chi-postgresql/pkg/validator"
)

func RegisterUserRoutes(r chi.Router, handler *handler.UserHandler) {
	r.Route("/auth", func(r chi.Router) {
		r.With(validator.ValidateRequest[dto.UserLoginRequest]()).Post("/login", handler.Login)
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", handler.GetAll)
		r.Get("/{id}", handler.GetById)

		r.Group(func(r chi.Router) {
			r.Use(middleware.JWTAuth)

			r.With(validator.ValidateRequest[dto.CreateUserRequest]()).Post("/", handler.Create)
			r.With(validator.ValidateRequest[dto.UpdateUserRequest]()).Put("/{id}", handler.Update)
			r.Delete("/{id}", handler.Delete)
		})
	})
}
