package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() {
	db, err := gorm.Open(sqlite.Open("./database/db.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{}, &Hosts{})
	log.Printf("[Resty] All Migrations Applied")
}
