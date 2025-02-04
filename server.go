package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
)

func router() *fiber.App {
	engine := handlebars.New("./views", ".hbs")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/portal/login")
	})

	portal := app.Group("/portal")
	portal.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{
			"title": "Login",
		}, "layouts/main")
	})

	return app
}

func main() {
	app := router()

	log.Fatal(app.Listen(":3000"))
}
