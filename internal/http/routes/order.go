package routes

import (
	"github.com/gofiber/fiber/v2"
	"oilio.com/internal/http/handlers"
)

func RegisterOrderRoutes(api fiber.Router) {
	r := api.Group("/order")
	r.Post("/", handlers.PlaceOrder())
}
