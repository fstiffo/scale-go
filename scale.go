package main

import (
	"time"

	"gorm.io/gorm"
	"sgajo.com/scale/scaledb"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := scaledb.Open()
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	scaledb.AutoMigrate(db)

	// Create
	michela := scaledb.Owner{
		Model:        gorm.Model{},
		Name:         "michels",
		DuesPayments: []scaledb.JournalEntry{},
	}
	db.Create(&michela)

	je := scaledb.NewStairsPayment(time.Date(2019, 7, 1, 0, 0, 0, 0, time.Local))
	db.Create(&je)

	// Read
	// var product Product
	// db.First(&product, 1)                 // find product with integer primary key
	// db.First(&product, "code = ?", "D42") // find product with code D42

	// // Update - update product's price to 200
	// db.Model(&product).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - delete Owner
	db.Delete(&michela, 1)
}
