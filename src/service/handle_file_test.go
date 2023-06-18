package service

import (
	"os"
	"strings"
	"testing"

	"github.com/mviniciusgc/desafio_neoway/src/mocks"
	"github.com/stretchr/testify/assert"
)

func setupService() *ClientService {
	as := InitializeService(&mocks.IRepository{})

	return as
}
func TestCreateArrayForInsert(t *testing.T) {

	t.Run("success in create array for insert", func(t *testing.T) {
		service := setupService()

		data, err := os.ReadFile("../mocks/file/base_teste.txt")
		if err != nil {
			assert.Nil(t, err)
		}

		r := strings.NewReader(string(data))

		resp, err := service.CreateArrayForInsert(r)
		assert.Nil(t, err)
		assert.NotNil(t, resp)

	})

	t.Run("error with read file", func(t *testing.T) {
		_, err := os.ReadFile("../mocks/file/sbase_teste.txt")
		if err != nil {
			assert.NotNil(t, err)
		}

	})

}
