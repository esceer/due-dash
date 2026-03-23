package setup

import (
	"net/http"

	"github.com/esceer/due-dash/backend/cmd/config"
	"github.com/esceer/due-dash/backend/internal/api"
	"github.com/esceer/due-dash/backend/internal/middleware"
	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"
)

func WebRouting(cfg *config.Config, services *serviceBundle) *echo.Echo {
	e := echo.New()

	// middlewares
	e.Use(middleware.Logger())
	e.Use(echo_middleware.CORSWithConfig(echo_middleware.CORSConfig{
		AllowOrigins:     []string{cfg.FrontendUrl},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	// e.Use(middleware.CSRF())  // TODO: configure csrfprotection (https://echo.labstack.com/docs/middleware/csrf)

	// system level api
	systemApi(e)

	// base path
	base := e.Group("/api/v1")

	// public routes
	publicApi(cfg, base, services)

	return e
}

func systemApi(e *echo.Echo) {
	e.Static("/spec", "api/spec")
	e.Static("/swagger-ui", "api/swagger-ui/dist")

	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
}

func publicApi(cfg *config.Config, g *echo.Group, services *serviceBundle) {
	healthApi := api.NewHealthApiHandler()
	health := g.Group("/health")
	health.GET("", healthApi.Health)

	templateApi := api.NewTemplateApiHandler(services.TemplateService)
	templates := g.Group("/templates")
	templates.GET("", templateApi.GetAll)
	templates.POST("", templateApi.Create)
	templates.GET("/:id", templateApi.GetById)
	templates.PUT("/:id", templateApi.Update)
	templates.DELETE("/:id", templateApi.Delete)

	taskApi := api.NewTaskApiHandler(services.TaskService)
	tasks := g.Group("/tasks")
	tasks.GET("", taskApi.GetAll)
	tasks.POST("", taskApi.Create)
	tasks.POST("generate", taskApi.GenerateFromTemplate)
	tasks.GET("/:id", taskApi.GetById)
	tasks.DELETE("/:id", taskApi.Delete)
	tasks.PATCH("/:id/status", taskApi.UpdateStatus)
}
