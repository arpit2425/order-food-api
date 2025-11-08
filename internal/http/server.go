package http

import (
	"github.com/gofiber/fiber/v2"
	"oilio.com/internal/http/routes"
	"oilio.com/internal/store/filestore"
)

func New() *fiber.App {
	app := fiber.New()
	app.Static("/docs", "./docs/openapi.yaml")
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok", "msg": "Server is up and running"})
	})
	store := filestore.New(
		"internal/store/filestore/data/products.json",
		"internal/store/filestore/data/orders.json",
		[]string{
			"internal/store/filestore/data/couponbase1.gz",
			"internal/store/filestore/data/couponbase2.gz",
			"internal/store/filestore/data/couponbase3.gz",
		},
	)
	routes.SetupRoutes(app, store)
	return app
}
