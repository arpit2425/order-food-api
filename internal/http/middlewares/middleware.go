package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"oilio.com/internal/config"
	"oilio.com/internal/http/helpers"
)

func Authenticate() fiber.Handler {
	expectedKey := config.Load().ApiKey

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
