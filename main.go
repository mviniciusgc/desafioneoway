package main

import (
	"github.com/mviniciusgc/desafio_neoway/src/server"
)

type mainService struct {
	Server server.Services
}

func main() {
	var s mainService
	s.Server.InitializeServer()
}
