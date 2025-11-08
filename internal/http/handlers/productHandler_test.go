package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"oilio.com/internal/model"
)

type MockStore struct {
	mock.Mock
}

var productList = []model.Product{{
	ID:       "1",
	Name:     "Waffle with Berries",
	Category: "Waffle",
}}

func (m *MockStore) ListProducts() ([]model.Product, error) {
	args := m.Called()
	return productList, args.Error(1)
}

func (m *MockStore) GetProduct(id string) (model.Product, error) {
	args := m.Called(id)
	return productList[0], args.Error(1)
}
func (m *MockStore) CreateOrder(model.Order) (model.Order, error) {
	return model.Order{}, nil

}
func (m *MockStore) ValidatePromo(code string) error {
	return nil
}

func setupFiber(handler fiber.Handler) (*fiber.App, *httptest.ResponseRecorder) {
	app := fiber.New()
	app.Get("/", handler)
	return app, httptest.NewRecorder()
}

func TestGetProducts_Success(t *testing.T) {
	mockStore := new(MockStore)
	expectedProducts := []interface{}{
		map[string]interface{}{"id": "1", "name": "Cake"},
		map[string]interface{}{"id": "2", "name": "Waffle"},
	}
	mockStore.On("ListProducts").Return(expectedProducts, nil)

	handler := NewProductHandler(mockStore)
	app := fiber.New()
	app.Get("/products", handler.GetProducts)

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockStore.AssertExpectations(t)
}

func TestGetProducts_Failure(t *testing.T) {
	mockStore := new(MockStore)
	mockStore.On("ListProducts").Return([]interface{}{}, errors.New("db error"))

	handler := NewProductHandler(mockStore)
	app := fiber.New()
	app.Get("/products", handler.GetProducts)

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	mockStore.AssertExpectations(t)
}

func TestGetProduct_Success(t *testing.T) {
	mockStore := new(MockStore)
	product := map[string]interface{}{"id": "10", "name": "Waffle"}
	mockStore.On("GetProduct", "10").Return(product, nil)

	handler := NewProductHandler(mockStore)
	app := fiber.New()
	app.Get("/products/:id", handler.GetProduct)

	req := httptest.NewRequest(http.MethodGet, "/products/10", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockStore.AssertExpectations(t)
}

func TestGetProduct_MissingID(t *testing.T) {
	mockStore := new(MockStore)
	handler := NewProductHandler(mockStore)
	app := fiber.New()
	app.Get("/products/:id?", handler.GetProduct)

	req := httptest.NewRequest(http.MethodGet, "/products/", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestGetProduct_NotFound(t *testing.T) {
	mockStore := new(MockStore)
	mockStore.On("GetProduct", "999").Return(nil, errors.New("not found"))

	handler := NewProductHandler(mockStore)
	app := fiber.New()
	app.Get("/products/:id", handler.GetProduct)

	req := httptest.NewRequest(http.MethodGet, "/products/999", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	mockStore.AssertExpectations(t)
}
