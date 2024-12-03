package main

import (
	// "net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// e.GET("/", serveFrontend) // Serve the frontend in production.
	// e.GET("/api/health", healthCheck)
	// e.GET("/api/users", getUsers)

	e.Logger.Fatal(e.Start(":8080"))
}