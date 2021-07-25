package scaledb

import (
	"github.com/uniplaces/carbon"
	"gorm.io/gorm"
)

func Cash(db *gorm.DB) int {
	var result struct {
		TotalDebits  int
		TotalCredits int
	}
	db.Model(&JournalEntry{}).
		Select("SUM(debit) as total_debits, SUM(credit) as total_credits").
		First(&result)
	return result.TotalCredits - result.TotalDebits
}

func ToBeRepaid(db *gorm.DB) int {
	var result struct {
		TotalDebits  int
		TotalCredits int
	}
	db.Model(&JournalEntry{}).
		Select("SUM(debit) as total_debits, SUM(credit) as total_credits").
		Where("account = ? OR account = ?", Loan, Repayment).
		First(&result)
	return result.TotalDebits - result.TotalCredits
}

func NetWorth(db *gorm.DB) int {
	var result struct {
		TotalDebits  int
		TotalCredits int
	}

	db.Model(&JournalEntry{}).
		Select("SUM(debit) as total_debits").
		Where("account = ?", StairsPayment).
		First(&result)
	totalStairsPayments := result.TotalDebits

	db.Model(&JournalEntry{}).
		Select("SUM(debit) as total_debits, SUM(credit) as total_credits").
		Where("account = ? OR account = ?", Expenditure, Revenue).
		First(&result)
	totalOthers := result.TotalCredits - result.TotalDebits

	var param Param
	db.Last(&param)
	// p := period.Between(param.ValidFrom, time.Now())
	today := carbon.Now()
	from := carbon.NewCarbon(param.ValidFrom)
	carbonMonths := today.DiffInMonths(from, true) - 24 // -24 because there is a bug in the carbon library!
	// months := p.Months() + p.Years()*12
	var numOfOwners int64
	db.Model(&Owner{}).Count(&numOfOwners)
	//mesi * num_condomini * (snd s.attuale).quotaMensile + altro - pagamenti
	return int(carbonMonths)*int(numOfOwners)*param.MonthlyDuesRate + totalOthers - totalStairsPayments
}

func AlreadyPayedRates(db *gorm.DB, ownerID uint) int {
	var result struct {
		TotalDebits  int
		TotalCredits int
	}

	db.Model(&JournalEntry{}).
		Select("SUM(credit) as total_credits").
		Where("owner_id = ?", ownerID).
		First(&result)
	totalDuesPayments := result.TotalCredits

	var param Param
	db.Last(&param)

	payedMonthlyRates := totalDuesPayments / param.MonthlyDuesRate

	today := carbon.Now()
	from := carbon.NewCarbon(param.ValidFrom)
	carbonMonths := today.DiffInMonths(from, true) - 24 // -24 because there is a bug in the carbon library!

	return payedMonthlyRates - int(carbonMonths)
}
