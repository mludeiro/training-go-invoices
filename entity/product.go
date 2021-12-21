package models

type Product struct {
	Id   uint   `gorm:"primaryKey"`
	Name string ``
}

func (Product) TableName() string {
	return "Product"
}
