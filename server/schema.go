package main

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Hosts struct {
	ID        uint `gorm:"primaryKey"`
	Hostname  string
	Proxy     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
