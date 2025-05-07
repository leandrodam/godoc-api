package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/leandrodam/godoc-api/internal/adapters/http/routes"
	"github.com/leandrodam/godoc-api/internal/config"
)

func main() {
	cfg := config.Load()

	e := echo.New()

	routes.RegisterRoutes(e)

	log.Printf("starting server on port %s", cfg.App.Port)

	if err := e.Start(":" + cfg.App.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
