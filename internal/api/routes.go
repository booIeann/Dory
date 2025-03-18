package api

import "github.com/gofiber/fiber/v2"

func SetupAPIRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/login", LoginHandler)
	api.Get("/logout", LogoutHandler)
}
