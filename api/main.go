package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/sagoresarker/url-shortner-go/routes"
)

func setupRoute(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	app := fiber.New()
	app.Use(logger.New())
	setupRoute(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
