package handlers

import (
	"github.com/gofiber/fiber/v2"
	"oilio.com/internal/http/helpers"
	"oilio.com/internal/store"
)

func GetProducts(store store.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		products, err := store.ListProducts()
		if err != nil {
			return helpers.InternalServerError(c, err.Error())
		}
		return helpers.Success(c, products)
	}
}
