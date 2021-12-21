package models

import (
	"time"
)

type Invoice struct {
	Id          uint      `gorm:"primaryKey"`
	ClientId    uint      `gorm:"index"`
	InvoiceDate time.Time ``
	Products    []Product `json:",omitempty"`
}

func (Invoice) TableName() string {
	return "Invoice"
}
