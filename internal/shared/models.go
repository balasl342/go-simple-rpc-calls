package shared

type Product struct {
	Id     int    `gorm:"primary_key;autoIncrement" json:"id"`
	Name   string `gorm:"column:product_name" json:"product_name"`
	ProdID int    `json:"product_id" gorm:"column:product_id"`
	Price  int    `json:"product_price" gorm:"column:product_price"`
}
