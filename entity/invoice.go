package entity

import (
	"time"
)

type Invoice struct {
	Id             uint             `gorm:"primaryKey"`
	ClientId       uint             `gorm:"index"`
	InvoiceDate    time.Time        ``
	InvoiceDetails []InvoiceDetails ``
}

func (Invoice) TableName() string {
	return "Invoice"
}
