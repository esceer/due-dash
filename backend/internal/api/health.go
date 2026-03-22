package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthApi interface {
	Health(c echo.Context) error
}

type healthApi struct{}

func NewHealthApiHandler() HealthApi {
	return &healthApi{}
}

func (a *healthApi) Health(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
