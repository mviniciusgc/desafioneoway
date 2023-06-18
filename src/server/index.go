package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mviniciusgc/desafio_neoway/src/db"
	"github.com/mviniciusgc/desafio_neoway/src/router"
	"github.com/spf13/viper"
)

type Services struct {
	HandlerServer router.HandlerServices
}

func (s Services) InitializeServer() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to initialize, error to read .env: %v", err)
	}

	d := db.InitializeClientBD()

	err = d.InitializeTables()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to initialize data base: %w", err))
	}

	r, err := s.HandlerServer.CreateRouterServices()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to initialize services: %w", err))
	}

	port := viper.GetString("API_PORT")

	fmt.Println("Server is running on port", port)
	err = http.ListenAndServe(port, r.Route)
	if err != nil {
		log.Fatal("Error initialize the server: ", err)
	}

}
