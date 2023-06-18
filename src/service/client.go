package service

import "github.com/mviniciusgc/desafio_neoway/src/repositorie"

type ClientService struct {
	clientrepository repositorie.IRepository
}

func InitializeService(r repositorie.IRepository) *ClientService {
	return &ClientService{
		clientrepository: r,
	}
}
