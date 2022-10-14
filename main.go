package main

import "github.com/snowcat471/simple-sns-user-service/server"

func main() {
	server := server.NewFiberServer(3000)
	server.Run()
}
