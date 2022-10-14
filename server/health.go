package server

import "github.com/gofiber/fiber/v2"

func health(c *fiber.Ctx) error {
	c.SendString("OK")
	return nil
}
