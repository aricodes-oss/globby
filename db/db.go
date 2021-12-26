package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Unable to connect to database")
	}

	db.AutoMigrate(&Lobby{})

	Connection = db
}
