package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	app  *fiber.App
	port int
}

func NewServer(port int) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandelr,
	})
	app.Use(logger.New())

	server := &Server{app, port}
	server.addRoutes()

	return server
}

func (s *Server) Run() {
	err := s.app.Listen(fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatalln(err)
	}
}

func (s *Server) addRoutes() {
	s.app.Get("/health", health)
}
