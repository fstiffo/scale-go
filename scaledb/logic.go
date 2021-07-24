package scaledb

import "gorm.io/gorm"

func Cash(db *gorm.DB) int {
	var result struct {
		TotalDebit  int
		TotalCredit int
	}
	db.Model(&JournalEntry{}).
		Select("SUM(debit) as total_debit, SUM(credit) as total_credit").
		First(&result)
	return result.TotalCredit - result.TotalDebit
}

func ToBeRepaid(db *gorm.DB) int {
	var result struct {
		TotalDebit  int
		TotalCredit int
	}
	db.Model(&JournalEntry{}).
		Select("SUM(debit) as total_debit, SUM(credit) as total_credit").
		Where("account = ? OR account = ?", Loan, Repayment).
		First(&result)
	return result.TotalDebit - result.TotalCredit
}
