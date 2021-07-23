package scaledb

import (
	"time"

	"gorm.io/gorm"
)

type Owner struct {
	gorm.Model
	Name         string
	DuesPayments []JournalEntry
}

const (
	StairsPayment = iota + 1
	Loan
	Repayment
	DuesPayment
	Expenditure
	Revenue
)

type JournalEntry struct {
	gorm.Model
	Date        time.Time
	Debit       int
	Credit      int
	Account     int
	OwnerID     uint
	Description string
}

func NewStairsPayment(date time.Time) *JournalEntry {
	je := JournalEntry{
		Model:   gorm.Model{},
		Date:    date,
		Debit:   20,
		Credit:  0,
		Account: StairsPayment,
	}
	return &je
}

func NewLoan(date time.Time, amount int) *JournalEntry {
	je := JournalEntry{
		Model:   gorm.Model{},
		Date:    date,
		Debit:   amount,
		Credit:  0,
		Account: Loan,
	}
	return &je
}

func NewRepayment(date time.Time, amount int) *JournalEntry {
	je := JournalEntry{
		Model:   gorm.Model{},
		Date:    date,
		Debit:   0,
		Credit:  amount,
		Account: Repayment,
	}
	return &je
}

func NewDuesPayment(date time.Time, amount int, ownerID uint) *JournalEntry {
	je := JournalEntry{
		Model:   gorm.Model{},
		Date:    date,
		Debit:   0,
		Credit:  amount,
		Account: DuesPayment,
		OwnerID: ownerID,
	}
	return &je
}

func NewExpenditure(date time.Time, amount int, desctiption string) *JournalEntry {
	je := JournalEntry{
		Model:       gorm.Model{},
		Date:        date,
		Debit:       amount,
		Credit:      0,
		Account:     Expenditure,
		Description: desctiption,
	}
	return &je
}

func NewRevenue(date time.Time, amount int, desctiption string) *JournalEntry {
	je := JournalEntry{
		Model:       gorm.Model{},
		Date:        date,
		Debit:       0,
		Credit:      amount,
		Account:     Revenue,
		Description: desctiption,
	}
	return &je
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Owner{})
	db.AutoMigrate(&JournalEntry{})
}
