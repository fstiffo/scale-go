package scaledb

import (
	"time"

	"gorm.io/gorm"
)

func CreateOwner(db *gorm.DB, name string) (*gorm.DB, *Owner) {
	owner := Owner{
		Model: gorm.Model{},
		Name:  name,
	}
	result := db.Create(&owner)
	return result, &owner
}

func CreateStairsPayment(db *gorm.DB, date time.Time) (*gorm.DB, *JournalEntry) {
	je := NewStairsPayment(date)
	result := db.Create(je)
	return result, je
}

func CreateLoan(db *gorm.DB, date time.Time, amount int) (*gorm.DB, *JournalEntry) {
	je := NewLoan(date, amount)
	result := db.Create(je)
	return result, je
}
