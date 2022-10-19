package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	app *fiber.App
}

func NewServer() *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})
	app.Use(cors.New())
	app.Use(logger.New())

	server := &Server{app}
	server.addRoutes()

	return server
}

func (s *Server) Run(port int) {
	if err := s.app.Listen(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalln(err)
	}
}

func (s *Server) addRoutes() {
	s.app.Get("/health", health)
}
