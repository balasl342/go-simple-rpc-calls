package rpc

import (
	"testing"

	"github.com/balasl342/go-simple-rpc-calls/internal/shared"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking the ProductRepository
type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) CreateProduct(product *shared.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) GetProductById(prod shared.Product, id int) (shared.Product, error) {
	args := m.Called(prod, id)
	return args.Get(0).(shared.Product), args.Error(1)
}

func TestCreateProducts(t *testing.T) {
	mockRepo := new(MockProductRepository)
	mockProduct := &shared.Product{Name: "TestProduct"}
	mockRepo.On("CreateProduct", mockProduct).Return(nil)

	service := RPCService{Repo: mockRepo}
	var reply shared.Product

	err := service.CreateProducts(mockProduct, &reply)
	assert.Nil(t, err)
	assert.Equal(t, *mockProduct, reply)

	mockRepo.AssertExpectations(t)
}

func TestGetProductsByID(t *testing.T) {
	mockRepo := new(MockProductRepository)
	mockProduct := shared.Product{ProdID: 1, Name: "TestProduct"}
	mockRepo.On("GetProductById", shared.Product{}, 1).Return(mockProduct, nil)

	service := RPCService{Repo: mockRepo}
	var reply shared.Product

	request := &shared.Product{ProdID: 1}
	err := service.GetProductsByID(request, &reply)
	assert.Nil(t, err)
	assert.Equal(t, mockProduct, reply)

	mockRepo.AssertExpectations(t)
}
