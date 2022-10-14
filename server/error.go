package server

import (
	"github.com/gofiber/fiber/v2"
)

func errorHandelr(c *fiber.Ctx, err error) error {

	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	res := NewErrorResponse(code, message)

	return c.Status(code).JSON(res)
}
