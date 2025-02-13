package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("../views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New(logger.Config{
		Format: "[${ip}] ${time} ${status} - ${method} ${path} ${ua} \n",
	}))

	PrepareMiddleware(app)

	log.Fatal(app.Listen(":3000"))
}
