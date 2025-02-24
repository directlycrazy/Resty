package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("./database/data/db.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{}, &Hosts{})
	Database = db
	log.Printf("[Resty] All Migrations Applied")
}
