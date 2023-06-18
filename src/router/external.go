package router

import (
	"net/http"

	"github.com/spf13/viper"
)

type Line struct {
	Values []string
}

func handle(s *HandlerServices) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Obtém o arquivo enviado com o nome de campo "file"
		file, _, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Defini o tamanho do lote que será inserido no banco
		lengthBatch := viper.GetInt("LENGTH_BATCH")

		err = s.ClientController.Handle(file, lengthBatch)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)

	})

}
