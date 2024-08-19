package repository

import (
	"fmt"

	"github.com/balasl342/go-simple-rpc-calls/internal/shared"
	"github.com/jinzhu/gorm"
)

type ProductRepository struct {
	Db *gorm.DB
}

func (repo *ProductRepository) CreateProduct(Product *shared.Product) error {
	var count int64
	if err := repo.Db.Table("products").Where("product_name=?", Product.Name).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("Product already exists")
	}
	err := repo.Db.Table("products").Create(&Product).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ProductRepository) GetProductById(prod shared.Product, id int) (shared.Product, error) {
	if err := repo.Db.Table("products").Where("product_id=?", id).Find(&prod).Error; err != nil {
		return shared.Product{}, err
	}
	return prod, nil

}
