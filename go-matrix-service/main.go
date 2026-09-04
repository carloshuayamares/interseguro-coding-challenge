package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/interseguro/matrix-service/client"
	"github.com/interseguro/matrix-service/handlers"
	"github.com/interseguro/matrix-service/middleware"
	"github.com/interseguro/matrix-service/services"
)

func main() {
	port := getenv("PORT", "8080")
	jwtSecret := getenv("JWT_SECRET", "change-me-in-production")
	nodeServiceURL := getenv("NODE_SERVICE_URL", "http://localhost:3000/api/v1/stats")

	httpClient := &client.HTTPClient{Client: &http.Client{Timeout: 10 * time.Second}, URL: nodeServiceURL}
	matrixService := services.NewMatrixService(httpClient)
	matrixHandler := handlers.NewMatrixHandler(matrixService)

	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(middleware.JWT(jwtSecret))
	app.Post("/api/v1/matrix/process", matrixHandler.Process)
	app.Post("/api/v1/matrix/rotate", matrixHandler.Rotate)

	log.Printf("matrix service listening on :%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
