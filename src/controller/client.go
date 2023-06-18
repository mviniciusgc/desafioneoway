package controller

import (
	"github.com/mviniciusgc/desafio_neoway/src/service"
)

type ClientController struct {
	ClientService service.IService
}

func InitializeController(s service.IService) *ClientController {
	return &ClientController{
		ClientService: s,
	}
}
