package main

type Product struct {
	Id     int    `gorm:"primary_key;autoIncrement" json:"id"`
	Name   string `gorm:"column:product_name" json:"product_name"`
	ProdID int    `json:"product_id" gorm:"column:product_id"`
	Price  int    `json:"product_price" gorm:"column:product_price"`
}

type Service struct{}

// post the countries in the Product table
func (s Service) CreateProducts(request *Product, reply *Product) error {
	err := CreateProduct(request)
	*reply = *request
	return err
}

func (s Service) GetProductsByID(request *Product, reply *Product) error {
	var coun Product

	Product, err := GetProductById(coun, request.ProdID)
	if err != nil {
		return err
	}
	*reply = Product
	return nil
}
