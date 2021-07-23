package scaledb

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Open() (*gorm.DB, error) {
	// github.com/mattn/go-sqlite3
	return gorm.Open(sqlite.Open("scale.db"), &gorm.Config{})
}
