package models

import (
	"time"
)

type Invoice struct {
	Id          uint      `gorm:"primaryKey"`
	ClientId    uint      `gorm:"index"`
	ClientName  string    `gorm:"not null;default:null"`
	InvoiceDate time.Time ``
	Products    []Product `json:",omitempty"`
}

func (Invoice) TableName() string {
	return "Invoice"
}
