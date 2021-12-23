package entity

type Product struct {
	Id      uint     `gorm:"primaryKey"`
	Name    string   ``
	Invoice *Invoice `json:",omitempty"`
}

func (Product) TableName() string {
	return "Product"
}
