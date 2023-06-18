package router

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/mviniciusgc/desafio_neoway/src/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupExternalTest() (*HandlerServices, *httptest.Server, *chi.Mux) {
	s, _ := (&HandlerServices{ClientController: &mocks.IController{}}).CreateRouterServices()

	r := chi.NewRouter()

	ts := httptest.NewServer(r)

	return s, ts, r
}
func TestHandle(t *testing.T) {
	t.Run("request success with valid file", func(t *testing.T) {
		service, httpServer, router := setupExternalTest()
		router.Post("/api/external/uploadfile", handle(service))
		defer httpServer.Close()

		data, err := os.ReadFile("../mocks/file/small_base_teste.txt")
		if err != nil {
			assert.Nil(t, err)
		}

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)

		// Cria uma nova parte no corpo
		part, _ := writer.CreateFormFile("file", "filename.txt")

		// Escreve o conteúdo da parte
		part.Write(data)

		// Finaliza o escritor multipart
		writer.Close()

		service.ClientController.(*mocks.IController).
			On("Handle", mock.Anything, mock.Anything).
			Return(nil).
			Once()

		resp, _ := http.Post(httpServer.URL+"/api/external/uploadfile", writer.FormDataContentType(), body)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

	})
	t.Run("request fail with invalid file", func(t *testing.T) {
		service, httpServer, router := setupExternalTest()
		router.Post("/api/external/uploadfile", handle(service))
		defer httpServer.Close()

		data, err := os.ReadFile("../mocks/file/small_base_teste.txt")
		if err != nil {
			assert.Nil(t, err)
		}

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)

		// Cria uma nova parte no corpo
		part, _ := writer.CreateFormFile("file", "filename.txt")

		// Escreve o conteúdo da parte
		part.Write(data)

		// Finaliza o escritor multipart
		writer.Close()

		service.ClientController.(*mocks.IController).
			On("Handle", mock.Anything, mock.Anything).
			Return(errors.New("File is invalid")).
			Once()

		resp, _ := http.Post(httpServer.URL+"/api/external/uploadfile", writer.FormDataContentType(), body)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	})
	t.Run("request fail when param file is invalid", func(t *testing.T) {
		service, httpServer, router := setupExternalTest()
		router.Post("/api/external/uploadfile", handle(service))
		defer httpServer.Close()

		data, err := os.ReadFile("../mocks/file/small_base_teste.txt")
		if err != nil {
			assert.Nil(t, err)
		}

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)

		// Cria uma nova parte no corpo
		part, _ := writer.CreateFormFile("files", "filename.txt")

		// Escreve o conteúdo da parte
		part.Write(data)

		// Finaliza o escritor multipart
		writer.Close()

		service.ClientController.(*mocks.IController).
			On("Handle", mock.Anything, mock.Anything).
			Return(errors.New("File is invalid")).
			Once()

		resp, _ := http.Post(httpServer.URL+"/api/external/uploadfile", writer.FormDataContentType(), body)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	})

}
