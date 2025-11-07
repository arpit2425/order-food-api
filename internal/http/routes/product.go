package routes

import (
	"github.com/gofiber/fiber/v2"
	"oilio.com/internal/http/handlers"
)

func RegisterProductRoutes(api fiber.Router) {
	r := api.Group("/product")
	r.Get("/", handlers.GetProducts())
	// r.Get("/:id", handlers.GetProduct())

}
