package repository

import (
	"training-go-invoices/database"
	"training-go-invoices/entity"
)

type Invoice struct {
	Database *database.Database
}

func (this *Invoice) GetAll() ([]entity.Invoice, error) {

	var invoices []entity.Invoice

	db := this.Database.GetDB()

	db.Find(&invoices)

	return invoices, nil
}

func (this *Invoice) Add(invoice entity.Invoice) *entity.Invoice {
	if this.Database.GetDB().Create(invoice).RowsAffected != 1 {
		return nil
	}
	return &invoice
}

func (this *Invoice) Get(id uint) (*entity.Invoice, error) {

	var invoice *entity.Invoice

	db := this.Database.GetDB()

	db.Where("id = ?", id).Find(&invoice, id)

	return invoice, nil
}
