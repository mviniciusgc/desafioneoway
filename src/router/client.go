package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/mviniciusgc/desafio_neoway/src/controller"
	"github.com/mviniciusgc/desafio_neoway/src/db"
	"github.com/mviniciusgc/desafio_neoway/src/repositorie"
	"github.com/mviniciusgc/desafio_neoway/src/service"
)

type HandlerServices struct {
	ClientController controller.IController
	Route            *chi.Mux
}

func (se HandlerServices) CreateRouterServices() (*HandlerServices, error) {

	// iniciando o banco de dados
	db := db.InitializeClientBD()
	//inicando o repositorio
	cr := repositorie.InitializeClienteRepository(db)
	// iniciando o service
	sr := service.InitializeService(cr)
	// iniciando o controller
	c := controller.InitializeController(sr)

	d := &HandlerServices{
		ClientController: c,
	}
	// iniciando as rotas
	r := Handlers(d)
	se.Route = r

	return &se, nil
}
