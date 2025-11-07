package http

import (
	"github.com/gofiber/fiber/v2"
	"oilio.com/internal/http/routes"
)

func New() *fiber.App {
	app := fiber.New()
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok", "msg": "Server is up and running"})
	})
	routes.SetupRoutes(app)
	return app
}
