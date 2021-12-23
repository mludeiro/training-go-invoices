package repository

import (
	"training-go-invoices/entity"

	"gorm.io/gorm"
)

//Is this ok?
var db = &gorm.DB{}

func GetAll() ([]entity.Invoice, error) {

	var invoices []entity.Invoice

	db.Find(&invoices)

	return invoices, nil
}

//TODO
func Add(product entity.Invoice) (uint, error) {
	return 0, nil
}

func Get(id uint) (*entity.Invoice, error) {

	var invoice *entity.Invoice

	db.Where("id = ?", id).Find(&invoice, id)

	return invoice, nil
}
