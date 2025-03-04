package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func Init() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Fatalf("%v", err)
			return RenderError(c, 500)
		},
		DisableStartupMessage: true,
	})

	app.Use(logger.New(logger.Config{
		Format: "[${ip}] ${time} ${status} - ${method} ${path} ${ua} \n",
	}))

	PrepareMiddleware(app)

	log.Printf("[Resty] Available at %s:%d", RestyConfig.Host, RestyConfig.Port)
	log.Fatal(app.Listen(fmt.Sprintf("%s:%d", RestyConfig.Host, RestyConfig.Port)))
}
