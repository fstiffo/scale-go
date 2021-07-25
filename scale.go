package main

import (
	"fmt"

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
	// scaledb.AutoMigrate(db)

	// Import from former project db
	// scaledb.Import(db)

	fmt.Println("Cash: ", scaledb.Cash(db))
	fmt.Println("To be repaid: ", scaledb.ToBeRepaid(db))
	fmt.Println("Net worth: ", scaledb.NetWorth(db))
	for id := uint(1); id < 5; id++ {
		var owner scaledb.Owner
		db.First(&owner, id)
		fmt.Println("Already payed monthly rates by "+owner.Name+": ", scaledb.AlreadyPayedRates(db, id))
	}

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
	// db.Delete(michela, 1)
}
