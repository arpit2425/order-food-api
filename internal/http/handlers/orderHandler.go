package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"oilio.com/internal/http/helpers"
	"oilio.com/internal/model"
	"oilio.com/internal/store"
	"oilio.com/internal/validator"
)

func PlaceOrder(store store.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var orderRequest model.OrderRequest
		if err := c.BodyParser(&orderRequest); err != nil {
			return helpers.BadRequest(c, "Invalid JSON body")
		}
		if err := validator.ValidateOrderRequest(orderRequest); err != nil {
			return helpers.BadRequest(c, err.Error())
		}
		subtotal := calculateSubtotal(orderRequest.Items)
		discount := calculateDiscount(orderRequest.PromoCode)
		total := subtotal - discount
		order, err := store.CreateOrder(model.Order{
			ID:        uuid.NewString(),
			Items:     orderRequest.Items,
			Subtotal:  subtotal,
			Discount:  discount,
			Total:     total,
			PromoCode: &orderRequest.PromoCode,
		})
		if err != nil {
			return helpers.InternalServerError(c, err.Error())
		}
		return helpers.Success(c, order)
	}
}

func calculateSubtotal(items []model.OrderItem) float64 {
	subtotal := 0.0
	for _, item := range items {
		subtotal += 2.0 * float64(item.Quantity) // TODO: Make it dynamic from the product price
	}
	return subtotal
}

func calculateDiscount(promoCode string) float64 {
	return 0.0
}
