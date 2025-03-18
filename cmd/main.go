package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/joho/godotenv"

	"dory/internal/api"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	app := fiber.New()

	api.SetupAPIRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("public/index.html")
	})

	app.Get("/chat", api.ValidateTokenMiddleware, func(c *fiber.Ctx) error {
		return c.SendFile("public/chat.html")
	})

	app.Get("/ws", websocket.New(api.HandleWebSocket))

	app.Listen(":8080")
}
