package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/leandrodam/godoc-api/internal/adapters/http/handlers"
)

func RegisterPatientRoutes(e *echo.Echo) {
	v1 := e.Group("/v1")

	v1.POST("/patients", handlers.CreatePatient)
	v1.PUT("/patients/:slug", handlers.UpdatePatient)
	v1.DELETE("/patients/:slug", handlers.DeletePatient)
	v1.GET("/patients", handlers.GetPatients)
	v1.GET("/patients/:slug", handlers.GetPatient)
}
