package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/leandrodam/godoc-api/internal/adapters/http/handlers"
)

func RegisterHealthCheckRoutes(e *echo.Echo) {
	v1 := e.Group("/v1")

	v1.GET("/alive", handlers.IsAlive)
}
