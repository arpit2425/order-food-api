package filestore

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"oilio.com/internal/model"
	"oilio.com/internal/store"
	"oilio.com/internal/store/dao"
	"oilio.com/internal/store/s3store"
)

type fileStore struct {
	productDAO  *dao.FileDAO[model.Product]
	orderDAO    *dao.FileDAO[model.Order]
	couponStore *s3store.CouponValidator
}

func New(productPath, orderPath string, couponPaths []string) store.Store {
	ctx := context.Background()

	bucket := "orderfoodonline-files"
	keys := []string{
		"couponbase1.gz",
		"couponbase2.gz",
		"couponbase3.gz",
	}

	validator, err := s3store.NewCouponValidator(ctx, bucket, keys)
	if err != nil {
		log.Fatal(err)
	}

	fs := &fileStore{
		productDAO:  &dao.FileDAO[model.Product]{FilePath: productPath},
		orderDAO:    &dao.FileDAO[model.Order]{FilePath: orderPath},
		couponStore: validator,
	}
	println("[INFO] Loading coupons from:", couponPaths)
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
	ctx := context.Background()
	err := fs.couponStore.ValidatePromo(ctx, code)
	fmt.Println(err)
	return err
}
