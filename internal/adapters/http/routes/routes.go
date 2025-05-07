package routes

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo) {
	RegisterHealthCheckRoutes(e)
	RegisterDoctorRoutes(e)
	RegisterPatientRoutes(e)
}
