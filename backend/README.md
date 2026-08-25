# Go Backend Template

A production-level, highly scalable Go backend template built using the **Echo v4** HTTP framework. It utilizes connection pooled **PostgreSQL** via `pgx/v5` and compiled type-safe database queries via **SQLC**, backed by **Redis** and structured logging via **Zap**.

## Tech Stack

- **HTTP Framework**: [Echo v4](https://echo.labstack.com/)
- **Database Client**: [pgx/v5](https://github.com/jackc/pgx) with connection pooling (`pgxpool`)
- **Query Generator**: [SQLC](https://sqlc.dev/) for compilation-time type-safe database methods
- **Database Migrations**: [Goose](https://github.com/pressly/goose)
- **Configuration Parsing**: [caarlos0/env/v11](https://github.com/caarlos0/env) with `.env` loading using `godotenv`
- **Structured Logging**: [Uber Zap](https://github.com/uber-go/zap)
- **Caching**: [go-redis/v9](https://github.com/redis/go-redis)
- **Validation**: [go-playground/validator/v10](https://github.com/go-playground/validator)
- **Hot-Reloading**: [Air](https://github.com/air-verse/air)
- **API UI Documentation**: [Scalar API Reference](https://github.com/MarceloPetrucio/go-scalar-api-reference)

---

## Directory Structure

```
├── cmd/
│   └── api/
│       └── main.go            # Entrypoint & Graceful Shutdown
├── internal/
│   ├── controllers/           # HTTP handlers
│   ├── db/                    # SQLC generated code
│   ├── dto/                   # Request/Response structures & validation
│   ├── logging/               # Structured logging using Zap
│   ├── middlewares/           # Custom middlewares (JWT, Custom logger)
│   └── utils/                 # Utilities (DB Pool, Redis, config, validator)
├── database/
│   ├── queries/               # SQLC raw queries
│   └── schema/                # Goose SQL migrations
├── docs/                      # Documentation mirror explaining files
└── tests/                     # HTTP and unit test cases
```

---

## Getting Started

### Local Setup

1. Make sure you have **PostgreSQL** and **Redis** running locally.
2. Create your `.env` file from the example:
   ```bash
   cp .env.example .env
   ```
3. Initialize the database schema using Goose:
   ```bash
   make migrate-up
   ```
4. Run the application with hot-reloading:
   ```bash
   make dev
   ```

### Docker Setup

1. Start all services (PostgreSQL, Redis, API) via docker compose:
   ```bash
   docker compose up --build
   ```

---

## Interactive Documentation

Once the backend is running, navigate to:
```
http://localhost:8080/docs
```
to view the interactive API reference served via Scalar.

---

## Makefile Automation

- `make build`: Compile the Go application to `bin/api`.
- `make run`: Run the server locally.
- `make dev`: Run the server with hot-reloading (`air`).
- `make test`: Run automated tests.
- `make lint`: Run code formatting and static analysis checkers (`golangci-lint`).
- `make vulncheck`: Scan dependencies for known security vulnerabilities (`govulncheck`).
- `make sqlc`: Compile SQL queries into Go code using SQLC.
- `make migrate-up`: Run Goose database migrations.
- `make migrate-down`: Rollback the last Goose migration.
- `make migrate-status`: Inspect migration status.
- `make docker-up`: Start PostgreSQL, Redis, and API containers in daemon mode.
- `make docker-down`: Shut down Docker containers and clean volumes.

---

## 🔄 Workflow & Code Quality

### Git Hooks (Commit Message Validation)
To enforce professional commit names (following the Conventional Commits spec), run the following command to bind the git hooks path:
```bash
git config core.hooksPath .githooks
```
*(On Linux/macOS, ensure the hook is executable: `chmod +x .githooks/commit-msg`)*

### CI/CD Pipeline
Every Pull Request and commit pushed to `main`/`master`/`dev` triggers a GitHub Actions pipeline validating:
- Proper Go code formatting (`gofmt`).
- Static analysis checks (`go vet`).
- Code styling & security linters (`golangci-lint`).
- Automated tests suite completion (`go test`).
- Go application compilation health (`go build`).
