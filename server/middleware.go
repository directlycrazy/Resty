package main

import (
	"log"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func prepareDevServer(app *fiber.App) {
	url, err := url.Parse("http://localhost:5173")
	if err != nil {
		log.Fatal(err)
	}

	app.Use(proxy.Balancer(proxy.Config{
		Servers: []string{
			url.String(),
		},
		Next: func(c *fiber.Ctx) bool {
			return len(c.Path()) >= 4 && c.Path()[:4] == "/api"
		},
	}))
}

func PrepareMiddleware(app *fiber.App) {
	if os.Getenv("ENV") == "dev" {
		prepareDevServer(app)
		return
	}

	app.Use(helmet.New())

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Server", "Resty")
		return c.Next()
	})

	app.Use(func(c *fiber.Ctx) error {
		return HandleRoutes(c)
	})
}
