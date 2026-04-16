package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"

	"github.com/rana-touseef11/go-chi-postgresql/internal/app"
	"github.com/rana-touseef11/go-chi-postgresql/internal/config"
	"github.com/rana-touseef11/go-chi-postgresql/internal/middleware"
	"github.com/rana-touseef11/go-chi-postgresql/pkg/database"

	_ "github.com/rana-touseef11/go-chi-postgresql/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Golang Chi PostgreSQL
// @version v1
// @BasePath  /api/v1
// @securityDefinitions.apikey BearerAuth
// @name Authorization
// @in header
func main() {
	// load Config
	var cfg = config.MustLoad()

	// laod logger
	// logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// load DB
	pool := database.PostgreSQL()

	// load router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(httprate.LimitByIP(59, time.Minute))

	api := chi.NewRouter()

	app.NewAppRouter(api, pool)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Salamun Alaikum"))
	})

	if cfg.ENV != "prod" {
		r.Get("/swagger/*", httpSwagger.WrapHandler)
	}
	r.Mount("/api/v1", api)

	// load server
	server_Setup(cfg, r)
}

func server_Setup(cfg *config.Config, r *chi.Mux) {
	const fiveSec = 5 * time.Second
	const tenSec = 2 * fiveSec
	server := &http.Server{
		Addr:         cfg.Addr(),
		Handler:      r,
		ReadTimeout:  fiveSec,
		WriteTimeout: tenSec,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		slog.Info("Server started...", slog.String("Addr", server.Addr))

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Failed to start server", slog.String("Error", err.Error()))
		}
	}()
	sig := <-done
	slog.Info("Signal received", slog.String("signal", sig.String()))

	ctx, cancel := context.WithTimeout(context.Background(), fiveSec)
	defer cancel()

	slog.Info("Shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server Failed to shutdown", slog.String("Error", err.Error()))
	}
}
