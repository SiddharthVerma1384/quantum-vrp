package controllers

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/CodeChefVIT/go-backend-template/internal/dto"
	"github.com/CodeChefVIT/go-backend-template/internal/utils"
	"github.com/labstack/echo/v4"
)

func HealthCheck(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Second)
	defer cancel()

	var pgStatus, redisStatus string
	var wg sync.WaitGroup

	wg.Add(2)

	// Check PostgreSQL
	go func() {
		defer wg.Done()
		if utils.DBPool == nil {
			pgStatus = "DOWN"
			return
		}
		if err := utils.DBPool.Ping(ctx); err != nil {
			pgStatus = "DOWN"
			return
		}
		pgStatus = "OK"
	}()

	// Check Redis
	go func() {
		defer wg.Done()
		if utils.RedisClient == nil {
			redisStatus = "DOWN"
			return
		}
		if err := utils.RedisClient.Ping(ctx).Err(); err != nil {
			redisStatus = "DOWN"
			return
		}
		redisStatus = "OK"
	}()

	wg.Wait()

	if pgStatus == "" {
		pgStatus = "DOWN"
	}
	if redisStatus == "" {
		redisStatus = "DOWN"
	}

	status := http.StatusOK
	if pgStatus == "DOWN" || redisStatus == "DOWN" {
		status = http.StatusServiceUnavailable
	}

	healthInfo := map[string]string{
		"postgres": pgStatus,
		"redis":    redisStatus,
	}

	return c.JSON(status, dto.NewSuccessResponse("Health check completed", healthInfo))
}
