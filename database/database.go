package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var D *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("./database/data/db.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{}, &Hosts{})
	D = db
	log.Printf("[Resty] All Migrations Applied")
}
