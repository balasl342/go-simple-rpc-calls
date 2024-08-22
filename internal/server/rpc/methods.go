package rpc

import (
	"github.com/balasl342/go-simple-rpc-calls/internal/server/repository"
	"github.com/balasl342/go-simple-rpc-calls/internal/shared"
)

type RPCService struct {
	Repo repository.ProductRepositoryInterface
}

func (r *RPCService) CreateProducts(request *shared.Product, reply *shared.Product) error {
	err := r.Repo.CreateProduct(request)
	if err != nil {
		return err
	}
	*reply = *request
	return nil
}

func (r *RPCService) GetProductsByID(request *shared.Product, reply *shared.Product) error {
	var coun shared.Product

	Product, err := r.Repo.GetProductById(coun, request.ProdID)
	if err != nil {
		return err
	}
	*reply = Product
	return nil
}
