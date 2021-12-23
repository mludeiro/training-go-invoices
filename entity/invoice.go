package entity

import (
	"time"
)

type Invoice struct {
	Id          uint      `gorm:"primaryKey"`
	ClientId    uint      `gorm:"index"`
	InvoiceDate time.Time ``
	Products    []uint    `json:",omitempty"`
}

func (Invoice) TableName() string {
	return "Invoice"
}
