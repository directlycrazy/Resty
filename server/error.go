package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func parseErrorCode(code uint16) (str string) {
	return http.StatusText(int(code))
}

func RenderError(c *fiber.Ctx, code uint16) (e error) {
	status := parseErrorCode(code)

	return c.Status(int(code)).Render("error", fiber.Map{
		"Code":   code,
		"Status": status,
	})
}
