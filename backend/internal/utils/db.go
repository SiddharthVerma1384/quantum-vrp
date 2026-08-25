package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/CodeChefVIT/go-backend-template/internal/logging"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DBPool *pgxpool.Pool

func InitDB() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		Config.PostgresUser,
		Config.PostgresPassword,
		Config.PostgresHost,
		Config.PostgresPort,
		Config.PostgresDB,
		Config.PostgresSSLMode,
	)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		logging.Fatalf("Unable to parse database DSN: %v", err)
	}

	config.MaxConns = Config.PostgresMaxConns
	config.MinConns = Config.PostgresMinConns
	config.MaxConnLifetime = Config.PostgresMaxConnLifetime
	config.MaxConnIdleTime = Config.PostgresMaxConnIdleTime

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	DBPool, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		logging.Fatalf("Failed to create database connection pool: %v", err)
	}

	if err := DBPool.Ping(ctx); err != nil {
		logging.Fatalf("Failed to ping database: %v", err)
	}

	logging.Infof("Database connection pool initialized successfully")
}

func CloseDB() {
	if DBPool != nil {
		DBPool.Close()
		logging.Infof("Database connection pool closed")
	}
}
