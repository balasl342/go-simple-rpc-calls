package rpc

import (
	"github.com/balasl342/go-simple-rpc-calls/internal/server/db"
	"github.com/balasl342/go-simple-rpc-calls/internal/server/repository"
	"github.com/balasl342/go-simple-rpc-calls/internal/shared"
)

type RPCService struct{}

func (*RPCService) CreateProducts(request *shared.Product, reply *shared.Product) error {
	userRepo := repository.ProductRepository{Db: db.Config.DB}
	err := userRepo.CreateProduct(request)
	*reply = *request
	return err
}

func (*RPCService) GetProductsByID(request *shared.Product, reply *shared.Product) error {
	var coun shared.Product

	userRepo := repository.ProductRepository{Db: db.Config.DB}
	Product, err := userRepo.GetProductById(coun, request.ProdID)
	if err != nil {
		return err
	}
	*reply = Product
	return nil
}
