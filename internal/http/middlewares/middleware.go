package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"oilio.com/internal/http/helpers"
)

func Authenticate() fiber.Handler {
	expectedKey := os.Getenv("API_KEY")
	if expectedKey == "" {
		expectedKey = "apitest"
	}

	return func(c *fiber.Ctx) error {
		apiKey := c.Get("api_key")

		if apiKey == "" {
			return helpers.Error(c, fiber.StatusUnauthorized, "Unauthorized")

		}

		if apiKey != expectedKey {
			return helpers.Error(c, fiber.StatusUnauthorized, "Unauthorized")
		}

		return c.Next()
	}
}
