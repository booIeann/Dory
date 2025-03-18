package api

import (
	"github.com/gofiber/fiber/v2"
)

func SuccessResponse(c *fiber.Ctx, status int, message string, data map[string]any) error {
	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"data":    data,
	})
}

func ErrorResponse(c *fiber.Ctx, status int, message string, details map[string]any) error {
	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"details": details,
	})
}
