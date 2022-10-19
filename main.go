package main

import (
	"github.com/snowcat471/simple-sns-user-service/config"
	"github.com/snowcat471/simple-sns-user-service/server"
)

func main() {
	s := server.NewServer()
	s.Run(config.ServerPort())
}
