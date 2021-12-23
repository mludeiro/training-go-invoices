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
	apple := entity.Product{Id: 1, Name: "Apple"}
	banana := entity.Product{Id: 2, Name: "Banana"}
	kiwi := entity.Product{Id: 3, Name: "Kiwi"}
	orange := entity.Product{Id: 4, Name: "Orange"}

	db.GetDB().Create(&apple).Create(&banana).Create(&kiwi).Create(&orange)

	productList1 := []entity.Product{apple, banana}
	productList2 := []entity.Product{apple, kiwi, orange}

	invoice1 := entity.Invoice{Id: 1, ClientId: 1, InvoiceDate: time.Now(), Products: productList1}
	invoice2 := entity.Invoice{Id: 2, ClientId: 2, InvoiceDate: time.Now(), Products: productList2}

	db.GetDB().Create(&invoice1).Create(&invoice2)

	return db
}

func (db *Database) Migrate() *Database {
	db.GetDB().AutoMigrate(&entity.Invoice{}, &entity.Product{})
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
