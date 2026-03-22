package middleware

import (
	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"
)

func Logger() echo.MiddlewareFunc {
	return echo_middleware.LoggerWithConfig(echo_middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","method":"${method}","status":"${status}","uri":"${uri}","error":"${error}"}` + "\n",
	})
}
