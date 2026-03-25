package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rana-touseef11/go-chi-postgresql/internal/config"
)

func PostgreSQL() *pgxpool.Pool {

	pool, err := pgxpool.New(context.Background(), config.MustLoad().DB_URL)

	if err != nil {
		log.Fatal("❌ Unable to connect: ", err)
	}

	// defer pool.Close()

	fmt.Println("✅ Connected to PostgreSQL DB!")

	return pool
}
