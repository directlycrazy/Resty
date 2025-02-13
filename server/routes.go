package main

import "github.com/gofiber/fiber/v2"

type Host struct {
	Fiber *fiber.App
}

var hosts = map[string]*Host{}

func CreateRoutes() {
	//Admin GUI
	admin := fiber.New()
	admin.Static("/", "../client/dist")
	admin.Use("*", func(c *fiber.Ctx) error {
		return c.SendFile("../client/dist/index.html")
	})
	hosts["localhost:3000"] = &Host{admin}
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
