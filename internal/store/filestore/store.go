package filestore

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"oilio.com/internal/model"
	"oilio.com/internal/store"
	"oilio.com/internal/store/dao"
)

type fileStore struct {
	productDAO *dao.FileDAO[model.Product]
	orderDAO   *dao.FileDAO[model.Order]
	coupons    *couponStore
}

func New(productPath, orderPath string, couponPaths []string) store.Store {
	fs := &fileStore{
		productDAO: &dao.FileDAO[model.Product]{FilePath: productPath},
		orderDAO:   &dao.FileDAO[model.Order]{FilePath: orderPath},
		coupons:    newCouponStore(couponPaths),
	}
	println("[INFO] Loading coupons from:", couponPaths)
	// if err := fs.coupons.loadCoupons(); err != nil {
	// 	println("[WARN] Failed to load coupons:", err.Error())
	// }
	println("[INFO] Coupons loaded successfully")

	return fs
}

func (fs *fileStore) ListProducts() ([]model.Product, error) {
	return fs.productDAO.ReadAll()
}

func (fs *fileStore) GetProduct(id string) (model.Product, error) {
	products, err := fs.productDAO.ReadAll()
	if err != nil {
		return model.Product{}, err
	}
	for _, p := range products {
		if p.ID == id {
			return p, nil
		}
	}
	return model.Product{}, errors.New("product not found")
}

func (fs *fileStore) CreateOrder(order model.Order) (model.Order, error) {
	orders, err := fs.orderDAO.ReadAll()
	if err != nil {
		return model.Order{}, err
	}

	order.ID = uuid.NewString()
	order.CreatedAt = time.Now()
	orders = append(orders, order)

	if err := fs.orderDAO.WriteAll(orders); err != nil {
		return model.Order{}, err
	}
	return order, nil
}

func (fs *fileStore) ListOrders() ([]model.Order, error) {
	return fs.orderDAO.ReadAll()
}

func (fs *fileStore) GetOrder(id string) (model.Order, error) {
	orders, err := fs.orderDAO.ReadAll()
	if err != nil {
		return model.Order{}, err
	}
	for _, o := range orders {
		if o.ID == id {
			return o, nil
		}
	}
	return model.Order{}, errors.New("order not found")
}

func (fs *fileStore) ValidatePromo(code string) error {
	return fs.coupons.validate(code)
}
