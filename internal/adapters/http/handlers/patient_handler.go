package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreatePatient(c echo.Context) error {
	return c.String(http.StatusCreated, "OK")
}

func UpdatePatient(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func DeletePatient(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func GetPatients(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func GetPatient(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
