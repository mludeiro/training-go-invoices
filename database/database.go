package database

import (
	"log"
	"time"
	"training-go-invoices/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	gormDB *gorm.DB
}

func (db *Database) CreateSampleData() *Database {

	productList1 := []uint{1, 2}
	productList2 := []uint{1, 3, 4}

	invoice1 := entity.Invoice{Id: 1, ClientId: 1, InvoiceDate: time.Now(), Products: productList1}
	invoice2 := entity.Invoice{Id: 2, ClientId: 2, InvoiceDate: time.Now(), Products: productList2}

	db.GetDB().Create(&invoice1).Create(&invoice2)

	return db
}

func (db *Database) Migrate() *Database {
	db.GetDB().AutoMigrate(&entity.Invoice{})
	return db
}

func (db *Database) GetDB() *gorm.DB {
	return db.gormDB
}

func (db *Database) InitializeSQLite() *Database {
	DB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"))

	if err != nil {
		log.Fatal("Cannot initialize database :(")
	}

	db.gormDB = DB
	return db
}
