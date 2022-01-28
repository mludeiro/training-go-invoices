package database

import (
	"log"
	"time"
	"training-go-invoices/entity"
	"training-go-invoices/tools"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	gormDB *gorm.DB
}

func (db *Database) CreateSampleData() *Database {

	invoice1 := entity.Invoice{Id: 1, ClientId: 1, InvoiceDate: time.Now()}
	invoice2 := entity.Invoice{Id: 2, ClientId: 2, InvoiceDate: time.Now()}

	db.GetDB().Create(&invoice1).Create(&invoice2)

	invoiceDetail1 := entity.InvoiceDetails{Id: 1, InvoiceId: invoice1.Id, ProductId: 1, Quantity: 2, UnitPrice: 500}
	invoiceDetail2 := entity.InvoiceDetails{Id: 2, InvoiceId: invoice1.Id, ProductId: 4, Quantity: 1, UnitPrice: 200}
	invoiceDetail3 := entity.InvoiceDetails{Id: 3, InvoiceId: invoice2.Id, ProductId: 8, Quantity: 5, UnitPrice: 100}

	db.GetDB().Create(&invoiceDetail1).Create(&invoiceDetail2).Create(&invoiceDetail3)

	return db
}

func (db *Database) Migrate() *Database {
	db.GetDB().AutoMigrate(&entity.Invoice{})
	db.GetDB().AutoMigrate(&entity.InvoiceDetails{})
	return db
}

func (db *Database) GetDB() *gorm.DB {
	return db.gormDB
}

func (db *Database) InitializeMySQL() *Database {
	dsn := "root:@tcp(127.0.0.1:3306)/goInvoices?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), createGormConfig())

	if err != nil {
		log.Fatal("Cannot initialize database :(")
	}

	db.gormDB = DB
	return db
}

func createGormConfig() *gorm.Config {
	return &gorm.Config{
		Logger: logger.New(
			tools.GetLogger(), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			}),
	}
}
