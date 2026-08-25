package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/CodeChefVIT/go-backend-template/internal/utils"
	_ "github.com/jackc/pgx/v5/stdlib" // SQL driver wrapper for pgx/v5
	"github.com/pressly/goose/v3"
)

const migrationsDir = "database/schema"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run cmd/migrate/main.go [up|down|status]")
	}

	command := os.Args[1]

	// Load environment configuration
	utils.LoadConfig()

	// Parse database DSN from config
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		utils.Config.PostgresUser,
		utils.Config.PostgresPassword,
		utils.Config.PostgresHost,
		utils.Config.PostgresPort,
		utils.Config.PostgresDB,
		utils.Config.PostgresSSLMode,
	)

	// Open connection using pgx stdlib compatibility layer
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer func() {
		_ = db.Close()
	}()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Set dialect to postgres
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Failed to set goose dialect: %v", err)
	}

	// Execute migrations
	switch command {
	case "up":
		log.Println("Executing migration command: up")
		if err := goose.Up(db, migrationsDir); err != nil {
			log.Fatalf("Goose migration UP failed: %v", err)
		}
	case "down":
		log.Println("Executing migration command: down")
		if err := goose.Down(db, migrationsDir); err != nil {
			log.Fatalf("Goose migration DOWN failed: %v", err)
		}
	case "status":
		log.Println("Executing migration command: status")
		if err := goose.Status(db, migrationsDir); err != nil {
			log.Fatalf("Goose migration STATUS failed: %v", err)
		}
	default:
		log.Fatal("Unknown migration command. Use 'up', 'down', or 'status'.")
	}

	log.Println("Migration command executed successfully")
}
