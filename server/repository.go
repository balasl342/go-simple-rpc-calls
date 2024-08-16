package main

import "fmt"

func CreateProduct(Product *Product) error {
	var count int64
	if err := Config.Db.Table("products").Where("product_name=?", Product.Name).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("Product already exists")
	}
	err := Config.Db.Table("products").Create(&Product).Error
	if err != nil {
		return err
	}
	return nil
}

func GetProductById(prod Product, id int) (Product, error) {
	if err := Config.Db.Table("products").Where("product_id=?", id).Find(&prod).Error; err != nil {
		return Product{}, err
	}
	return prod, nil

}
