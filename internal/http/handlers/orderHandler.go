package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"oilio.com/internal/http/helpers"
	"oilio.com/internal/model"
	"oilio.com/internal/store"
	"oilio.com/internal/validator"
)

type OrderHandler interface {
	PlaceOrder(c *fiber.Ctx) error
}

type orderHandler struct {
	store store.Store
}

func NewOrderHandler(store store.Store) OrderHandler {
	return &orderHandler{store: store}
}

func (h *orderHandler) PlaceOrder(c *fiber.Ctx) error {
	var orderRequest model.OrderRequest

	if err := c.BodyParser(&orderRequest); err != nil {
		return helpers.BadRequest(c, "Invalid JSON body")
	}

	if err := validator.ValidateOrderRequest(orderRequest); err != nil {
		return helpers.BadRequest(c, err.Error())
	}

	subtotal := 0.0
	for _, item := range orderRequest.Items {
		product, err := h.store.GetProduct(item.ProductID)
		if err != nil {
			return helpers.BadRequest(c, "Invalid product ID: "+item.ProductID)
		}
		subtotal += product.Price * float64(item.Quantity)
	}

	discount := h.calculateDiscount(orderRequest.PromoCode)
	total := subtotal - discount
	if total < 0 {
		total = 0
	}

	order := model.Order{
		ID:        uuid.NewString(),
		Items:     orderRequest.Items,
		Subtotal:  subtotal,
		Discount:  discount,
		Total:     total,
		PromoCode: &orderRequest.PromoCode,
	}

	savedOrder, err := h.store.CreateOrder(order)
	if err != nil {
		return helpers.InternalServerError(c, err.Error())
	}

	return helpers.Success(c, fiber.Map{
		"message": "Order placed successfully",
		"order":   savedOrder,
	})
}

func (h orderHandler) calculateDiscount(promoCode string) float64 {
	err := h.store.ValidatePromo(promoCode)
	if err != nil {
		return 0.0
	}
	return 2.0
}
