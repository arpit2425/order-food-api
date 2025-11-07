package handlers

import (
	"github.com/gofiber/fiber/v2"
	"oilio.com/internal/http/helpers"
	"oilio.com/internal/model"
	"oilio.com/internal/validator"
)

func PlaceOrder() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var orderRequest model.OrderRequest
		if err := c.BodyParser(&orderRequest); err != nil {
			return helpers.BadRequest(c, "Invalid JSON body")
		}
		if err := validator.ValidateOrderRequest(orderRequest); err != nil {
			return helpers.BadRequest(c, err.Error())
		}
		return helpers.Success(c, fiber.Map{"message": "Order placed successfully"})
	}
}
