package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rana-touseef11/go-chi-postgresql/internal/handler"
	"github.com/rana-touseef11/go-chi-postgresql/internal/repository"
	"github.com/rana-touseef11/go-chi-postgresql/internal/router"
	"github.com/rana-touseef11/go-chi-postgresql/internal/service"
)

type App struct {
	UserHandler *handler.UserHandler
	// OrderHandler *handler.OrderHandler
}

func NewApp(db *pgxpool.Pool) *App {
	// user
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// order
	// orderRepo := repository.NewOrderRepository(db)
	// orderService := service.NewOrderService(orderRepo)
	// orderHandler := handler.NewOrderHandler(orderService)

	return &App{
		UserHandler: userHandler,
		// OrderHandler: orderHandler,
	}
}

func NewAppRouter(r chi.Router, pool *pgxpool.Pool) {
	handle := NewApp(pool)

	router.RegisterUserRoutes(r, handle.UserHandler)
}
