package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type FiberServer struct {
	app  *fiber.App
	port int
}

func NewFiberServer(port int) Server {
	app := fiber.New()
	app.Use(logger.New())

	return &FiberServer{app, port}
}

func (s *FiberServer) Run() {
	s.app.Listen(fmt.Sprintf(":%d", s.port))
}
