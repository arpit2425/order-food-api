package routes

import "github.com/gofiber/fiber/v2"

func RegisterProductRoutes(api fiber.Router) {
	r := api.Group("/products")
	r.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok", "msg": "Get All product routes"})
	})
	// r.Get("/:id", handlers.GetProduct())

}
