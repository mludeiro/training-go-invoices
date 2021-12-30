package entity

type InvoiceDetails struct {
	Id        uint     `gorm:"primarykey"`
	InvoiceId uint     `gorm:"index"`
	Invoice   *Invoice `json:",omitempty"`
	ProductId uint     `gorm:"index"`
	Quantity  uint     `gorm:"not null"`
	UnitPrice uint     `gorm:"not null"`
}

func (InvoiceDetails) TableName() string {
	return "InvoiceDetails"
}
