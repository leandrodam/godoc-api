package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/leandrodam/godoc-api/internal/adapters/http/handlers"
)

func RegisterDoctorRoutes(e *echo.Echo) {
	v1 := e.Group("/v1")

	v1.POST("/doctors", handlers.CreateDoctor)
	v1.PUT("/doctors/:slug", handlers.UpdateDoctor)
	v1.DELETE("/doctors/:slug", handlers.DeleteDoctor)
	v1.GET("/doctors", handlers.GetDoctors)
	v1.GET("/doctors/:slug", handlers.GetDoctor)
}
