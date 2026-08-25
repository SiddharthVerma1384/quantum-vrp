# Makefile Commands Reference

## Overview

This document describes all available `make` commands in this Go backend template.

---

## Core Commands

| Command | Description |
|---------|-------------|
| `make` / `make all` | Default target. Runs `make build`. |
| `make build` | Compiles the Go binary to `bin/api` from `cmd/api/main.go`. |

---

## Development Commands

| Command | Description |
|---------|-------------|
| `make dev` | Starts hot-reload development server using **Air** (config: `.air.toml`). Auto-reloads on file changes. |
| `make run` | Runs the application directly with `go run cmd/api/main.go` (no hot-reload). |
| `make setup` | Configures Git hooks to use `.githooks` directory. Runs pre-commit checks automatically on `git commit`. Run once after cloning. |

---

## Testing & Code Quality

| Command | Description |
|---------|-------------|
| `make test` | Runs all tests with verbose output: `go test -v ./...` |
| `make lint` | Runs `golangci-lint` with config from `.golangci.yml` (linting, style, static analysis). |
| `make fmt` | Formats Go code using `go fmt ./...` (standard Go formatter). |
| `make vulncheck` | Runs `govulncheck ./...` to scan for known vulnerabilities in dependencies. |

---

## Database (SQLC & Migrations)

| Command | Description |
|---------|-------------|
| `make sqlc` | Generates Go code from SQL queries using `sqlc generate` (reads `sqlc.yaml`). Run after modifying SQL queries. |
| `make migrate-up` | Applies all pending migrations: `go run cmd/migrate/main.go up` |
| `make migrate-down` | Rolls back the last migration: `go run cmd/migrate/main.go down` |
| `make migrate-status` | Shows migration status (applied/pending): `go run cmd/migrate/main.go status` |

---

## Docker

| Command | Description |
|---------|-------------|
| `make docker-up` | Starts all Docker services in detached mode with rebuild: `docker compose up --build -d` |
| `make docker-down` | Stops and removes containers, networks, **and volumes**: `docker compose down -v` |

---

## Configuration

| Line | Purpose |
|------|---------|
| `-include .env` | Loads `.env` file (if present) and exports variables as environment variables. |
| `export` | Exports all variables to child processes. |

---

## Quick Reference

| Task | Command |
|------|---------|
| Daily development | `make dev` |
| Build for deployment | `make build` |
| Quick test run | `make run` |
| Run tests | `make test` |
| Lint before commit | `make lint` |
| Format code | `make fmt` |
| Security scan | `make vulncheck` |
| Start dependencies (DB, Redis, etc.) | `make docker-up` |
| Stop & clean Docker | `make docker-down` |
| Generate SQL code | `make sqlc` |
| Apply migrations | `make migrate-up` |
| Rollback migration | `make migrate-down` |
| Check migration status | `make migrate-status` |
| First-time setup | `make setup` |

---

## Prerequisites

Ensure these tools are installed:

| Tool | Install Command |
|------|-----------------|
| Go | `go install` |
| Air (hot reload) | `go install github.com/air-verse/air@latest` |
| golangci-lint | `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest` |
| govulncheck | `go install golang.org/x/vuln/cmd/govulncheck@latest` |
| sqlc | `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest` |
| migrate (custom) | `go build -o bin/migrate cmd/migrate/main.go` |
| Docker & Docker Compose | [Docker Desktop](https://www.docker.com/products/docker-desktop/) |

---

## Environment Setup

1. Copy `.env.example` to `.env` (if exists) and configure values
2. Run `make setup` to configure Git hooks
3. Run `make docker-up` to start dependencies
4. Run `make migrate-up` to apply migrations
5. Run `make dev` to start development
