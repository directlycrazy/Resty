package main

import (
	"github.com/directlycrazy/Resty/database"
	"github.com/directlycrazy/Resty/server"
)

func main() {
	database.InitDB()
	server.Init()
}
