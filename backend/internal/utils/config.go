package utils

import (
	"log"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type cfg struct {
	Env                     string        `env:"ENV" envDefault:"development"`
	Port                    string        `env:"PORT" envDefault:"8080"`
	JWTSecret               string        `env:"JWT_SECRET,notEmpty"`
	PostgresHost            string        `env:"POSTGRES_HOST,notEmpty"`
	PostgresPort            string        `env:"POSTGRES_PORT,notEmpty"`
	PostgresUser            string        `env:"POSTGRES_USER,notEmpty"`
	PostgresPassword        string        `env:"POSTGRES_PASSWORD,notEmpty"`
	PostgresDB              string        `env:"POSTGRES_DB,notEmpty"`
	PostgresSSLMode         string        `env:"POSTGRES_SSLMODE" envDefault:"disable"`
	PostgresMaxConns        int32         `env:"POSTGRES_MAX_CONNS" envDefault:"25"`
	PostgresMinConns        int32         `env:"POSTGRES_MIN_CONNS" envDefault:"5"`
	PostgresMaxConnLifetime time.Duration `env:"POSTGRES_MAX_CONN_LIFETIME" envDefault:"30m"`
	PostgresMaxConnIdleTime time.Duration `env:"POSTGRES_MAX_CONN_IDLE_TIME" envDefault:"15m"`
	RedisHost               string        `env:"REDIS_HOST" envDefault:"localhost"`
	RedisPort               string        `env:"REDIS_PORT" envDefault:"6379"`
	RedisPassword           string        `env:"REDIS_PASSWORD"`
	RedisPoolSize           int           `env:"REDIS_POOL_SIZE" envDefault:"20"`
	RedisMinIdleConns       int           `env:"REDIS_MIN_IDLE_CONNS" envDefault:"5"`
	FrontendURL             string        `env:"FRONTEND_URL" envDefault:"http://localhost:3000"`
	AdminURL                string        `env:"ADMIN_URL" envDefault:"http://localhost:3067"`
	RateLimitRPS            int           `env:"RATE_LIMIT_RPS" envDefault:"20"`
	ShutdownTimeout         time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"10s"`
	CORSOrigins             []string      // Computed field
}

var Config cfg

func LoadConfig() {
	// Load .env file if it exists; ignore error if it doesn't
	_ = godotenv.Load()

	if err := env.Parse(&Config); err != nil {
		log.Fatalf("Error parsing configuration: %v", err)
	}

	// Compute allowed CORS origins
	if Config.FrontendURL != "" {
		Config.CORSOrigins = append(Config.CORSOrigins, Config.FrontendURL)
	}
	if Config.AdminURL != "" {
		Config.CORSOrigins = append(Config.CORSOrigins, Config.AdminURL)
	}
}
