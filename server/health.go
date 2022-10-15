package server

import "github.com/gofiber/fiber/v2"

func health(c *fiber.Ctx) error {
	return c.SendString("OK")
}
