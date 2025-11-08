package handlers

import (
	"github.com/gofiber/fiber/v2"
	"oilio.com/internal/http/helpers"
	"oilio.com/internal/store"
)

type ProductHandler interface {
	GetProducts(c *fiber.Ctx) error
	GetProduct(c *fiber.Ctx) error
}

type productHandler struct {
	store store.Store
}

func NewProductHandler(store store.Store) ProductHandler {
	return &productHandler{store: store}
}
func (h *productHandler) GetProducts(c *fiber.Ctx) error {
	products, err := h.store.ListProducts()
	if err != nil {
		return helpers.InternalServerError(c, err.Error())
	}
	return helpers.Success(c, products)
}
func (h *productHandler) GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return helpers.BadRequest(c, "Product ID is required")
	}
	product, err := h.store.GetProduct(id)
	if err != nil {
		return helpers.NotFound(c, "Product not found")
	}
	return helpers.Success(c, product)
}
