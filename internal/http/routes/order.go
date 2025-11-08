package routes

import (
	"github.com/gofiber/fiber/v2"
	"oilio.com/internal/http/handlers"
	"oilio.com/internal/store"
)

func RegisterOrderRoutes(api fiber.Router, store store.Store) {
	r := api.Group("/order")
	handler := handlers.NewOrderHandler(store)
	r.Post("/", handler.PlaceOrder)
}
