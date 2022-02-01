package utils

import "github.com/gofiber/fiber/v2"

func ErrorResponse(errorMessage string) *fiber.Map {
	return &fiber.Map{"error": errorMessage}
}

func Response(c *fiber.Ctx, statusCode int, body *fiber.Map) error {
	c.Status(400)
	c.Type("application/json")
	return c.JSON(body)
}
