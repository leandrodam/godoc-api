package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func IsAlive(c echo.Context) error {
	return c.String(http.StatusCreated, "OK")
}
