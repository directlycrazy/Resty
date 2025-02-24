package server

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

type Host struct {
	Fiber *fiber.App
}

var hosts = map[string]*Host{}

func defineRoute(host string, proxyStr string) (e error) {
	r := fiber.New()

	proxyUrl, err := url.Parse(proxyStr)
	if err != nil {
		log.Fatal(err)
		return err
	}

	r.Use(proxy.Balancer(proxy.Config{
		Servers: []string{
			proxyUrl.String(),
		},
		ModifyResponse: func(c *fiber.Ctx) error {
			if RestyConfig.ServerHeaderEnabled {
				c.Response().Header.Add("Server", "resty")
			}
			return nil
		},
		Timeout: 30 * time.Second,
	}))

	hosts[fmt.Sprintf("%s:%v", host, RestyConfig.Port)] = &Host{r}

	return nil
}

func CreateRoutes() {
	//Admin GUI
	admin := fiber.New()
	admin.Use(compress.New())
	admin.Static("/", "./client/dist")
	admin.Use("*", func(c *fiber.Ctx) error {
		return c.SendFile("./client/dist/index.html")
	})
	hosts[fmt.Sprintf("%s:%v", RestyConfig.Host, RestyConfig.Port)] = &Host{admin}
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
