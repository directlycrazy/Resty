package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

type Host struct {
	Fiber *fiber.App
}

var hosts = map[string]*Host{}

func CreateRoutes() {
	//Admin GUI
	admin := fiber.New()
	admin.Use(compress.New())
	admin.Static("/", "../client/dist")
	admin.Use("*", func(c *fiber.Ctx) error {
		return c.SendFile("../client/dist/index.html")
	})
	hosts[fmt.Sprintf("%s:%v", RestyHost, RestyPort)] = &Host{admin}
}

func HandleRoutes(c *fiber.Ctx) (e error) {
	if len(hosts) == 0 {
		CreateRoutes()
	}

	host := hosts[c.Hostname()]
	if host == nil {
		return RenderError(c, 404)
	} else {
		host.Fiber.Handler()(c.Context())
		return nil
	}
}
