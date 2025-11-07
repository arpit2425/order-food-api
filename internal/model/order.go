package model

import "time"

type OrderItem struct {
	ProductID string `json:"productId" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required,min=1"`
}

type OrderRequest struct {
	PromoCode string      `json:"couponCode" validate:"omitempty,min=8,max=10"`
	Items     []OrderItem `json:"items" validate:"required,dive,required"`
}

type Order struct {
	ID        string      `json:"id"`
	Items     []OrderItem `json:"items"`
	Subtotal  float64     `json:"subtotal"`
	Discount  float64     `json:"discount"`
	Total     float64     `json:"total"`
	PromoCode *string     `json:"promoCode,omitempty"`
	CreatedAt time.Time   `json:"createdAt"`
}
