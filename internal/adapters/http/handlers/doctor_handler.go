package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateDoctor(c echo.Context) error {
	return c.String(http.StatusCreated, "OK")
}

func UpdateDoctor(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func DeleteDoctor(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func GetDoctors(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func GetDoctor(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
