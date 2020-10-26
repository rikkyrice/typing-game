package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// /healthcheck
// HealthCheckHandler ヘルスチェックするAPI
type HealthCheckHandler interface {
	HealthCheck() echo.HandlerFunc
}

type healthCheckHandler struct{}

func NewHealthCheckHandler() HealthCheckHandler {
	return &healthCheckHandler{}
}

func (h *healthCheckHandler) HealthCheck() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "HealthCheck OK")
	}
}
