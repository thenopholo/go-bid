package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/thenopholo/go-bid/internal/config"
	"github.com/thenopholo/go-bid/internal/server"
)

func main() {
	logger := config.NewLogger("MAIN_API")
	ctx := context.Background()
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	pool, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	))
	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	logger.Info("Starting application...")
	InitServer(pool)
}

func InitServer(pool *pgxpool.Pool) {
	server.NewServer(pool).Start()
}
