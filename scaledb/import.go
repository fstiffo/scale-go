package scaledb

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Import(db *gorm.DB) {
	// Create original param
	param := &Param{
		Model:             gorm.Model{},
		ValidFrom:         time.Date(2019, 7, 1, 0, 0, 0, 0, time.UTC),
		StairsCleaningFee: 20,
		CleaningsPerMonth: 2,
		MonthlyDues:       12,
	}
	db.Create(param)

	// Create original owners
	michela := &Owner{Name: "Michela"}
	db.Create(michela)
	gerardo := &Owner{Name: "Gerardo"}
	db.Create(gerardo)
	elena := &Owner{Name: "Elena"}
	db.Create(elena)
	giulia := &Owner{Name: "Giulia"}
	db.Create(giulia)

	statoscaledb, err := gorm.Open(sqlite.Open("stato_scale.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	// Raw SQL
	rows, err := statoscaledb.Raw("SELECT data, importo, type, condomino_id, causale FROM operazioni").Rows()

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			data         string
			importo      int
			_type        string
			condomino_id int
			causale      string
		)
		rows.Scan(&data, &importo, &_type, &condomino_id, &causale)
		date, _ := time.Parse(time.RFC3339, data)
		switch _type {
		case "pagamento_scale":
			je := NewStairsPayment(date)
			db.Create(je)
		case "prestito":
			je := NewLoan(date, importo)
			db.Create(je)
		case "restituzione":
			je := NewRepayment(date, importo)
			db.Create(je)
		case "versamento_quote":
			je := NewDuesPayment(date, importo, uint(condomino_id+1))
			db.Create(je)
		case "altra_spesa":
			je := NewExpenditure(date, importo, causale)
			db.Create(je)
		case "altro_versamento":
			je := NewRevenue(date, importo, causale)
			db.Create(je)
		default:

		}
	}
}
