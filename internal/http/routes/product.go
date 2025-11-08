package routes

import (
	"github.com/gofiber/fiber/v2"
	"oilio.com/internal/http/handlers"
	"oilio.com/internal/store"
)

func RegisterProductRoutes(api fiber.Router, store store.Store) {
	r := api.Group("/product")
	handler := handlers.NewProductHandler(store)
	r.Get("/", handler.GetProducts)
	r.Get("/:id", handler.GetProduct)

}
