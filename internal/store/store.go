package store

import "oilio.com/internal/model"

type Store interface {
	ListProducts() ([]model.Product, error)
	GetProduct(id string) (model.Product, error)

	CreateOrder(order model.Order) (model.Order, error)
}
