package api

import (
	"dory/internal/auth"
	"log"

	"github.com/gofiber/fiber/v2"
)

func ValidateTokenMiddleware(c *fiber.Ctx) error {
	token := c.Cookies("jwt")
	if token == "" {
		return c.Redirect("/", fiber.StatusFound)
	}

	claims, err := auth.ValidateToken(token)
	if err != nil {
		log.Println(err)
		return ErrorResponse(c, fiber.StatusUnauthorized, "Invalid or expired token", nil)
	}

	username := claims["username"].(string)
	c.Locals("username", username)

	return c.Next()
}
