package controller

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/mviniciusgc/desafio_neoway/src/entity"
	"github.com/mviniciusgc/desafio_neoway/src/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupController() *ClientController {
	as := InitializeController(&mocks.IService{})

	return as
}
func TestHandle(t *testing.T) {

	t.Run("success in formatting data to save in db", func(t *testing.T) {
		service := setupController()
		lengthBatch := 5

		newPurchase := []string{
			"064.532.359-43 1 0 2013-10-05 81,46 62,70 79.379.491/0001-83 79.379.491/0001-83",
			"598.756.230-91 0 0 2013-08-17 293,90 293,90 79.379.491/0001-83 79.379.491/0001-83",
		}

		CPFCNPJ := "123.456.789-00"
		Private := bool(true)
		Incomplete := bool(false)
		DateLastCompra := "2016-01-02"
		AverageTicket := float64(123.45)
		LastPurchaseTicket := float64(123.45)
		StoreMoreFrequent := "79.379.491/0001-83"
		StoreLastPurchase := "79.379.491/0001-83"

		purchasesFormated := []entity.Purchase{
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
		}
		service.ClientService.(*mocks.IService).
			On("CreateArrayForInsert", mock.Anything).
			Return(&newPurchase, nil).
			Once()

		service.ClientService.(*mocks.IService).
			On("ValidationData", mock.Anything).
			Return(nil).
			Once()

		service.ClientService.(*mocks.IService).
			On("FormateDataForCreate", mock.Anything).
			Return(&purchasesFormated).
			Once()

		service.ClientService.(*mocks.IService).
			On("InsertBatch", mock.Anything, mock.Anything).
			Return(nil).
			Once()

		data, _ := os.ReadFile("../mocks/file/base_teste.txt")

		r := strings.NewReader(string(data))

		err := service.Handle(r, lengthBatch)
		assert.Nil(t, err)

	})
	t.Run("error to create array for insert", func(t *testing.T) {
		service := setupController()
		lengthBatch := 5

		service.ClientService.(*mocks.IService).
			On("CreateArrayForInsert", mock.Anything).
			Return(nil, errors.New("error to ReadAll")).
			Once()

		data, _ := os.ReadFile("../mocks/file/small_base_teste.txt")

		r := strings.NewReader(string(data))

		err := service.Handle(r, lengthBatch)
		assert.NotNil(t, err)

	})

	t.Run("error to validate data", func(t *testing.T) {
		service := setupController()
		lengthBatch := 5

		newPurchase := []string{
			"064.532.359-43 1 0 2013-10-05 81,46 62,70 79.379.491/0001-83 79.379.491/0001-83",
			"598.756.230-91 0 0 2013-08-17 293,90 293,90 79.379.491/0001-83 79.379.491/0001-83",
		}

		service.ClientService.(*mocks.IService).
			On("CreateArrayForInsert", mock.Anything).
			Return(&newPurchase, nil).
			Once()

		service.ClientService.(*mocks.IService).
			On("ValidationData", mock.Anything).
			Return(errors.New("CPF or CNPJ invalid")).
			Once()

		data, _ := os.ReadFile("../mocks/file/base_teste.txt")

		r := strings.NewReader(string(data))

		err := service.Handle(r, lengthBatch)
		assert.NotNil(t, err)

	})

	t.Run("error to insert batch", func(t *testing.T) {
		service := setupController()
		lengthBatch := 5

		newPurchase := []string{
			"064.532.359-43 1 0 2013-10-05 81,46 62,70 79.379.491/0001-83 79.379.491/0001-83",
			"598.756.230-91 0 0 2013-08-17 293,90 293,90 79.379.491/0001-83 79.379.491/0001-83",
		}

		CPFCNPJ := "123.456.789-00"
		Private := bool(true)
		Incomplete := bool(false)
		DateLastCompra := "2016-01-02"
		AverageTicket := float64(123.45)
		LastPurchaseTicket := float64(123.45)
		StoreMoreFrequent := "79.379.491/0001-83"
		StoreLastPurchase := "79.379.491/0001-83"

		purchasesFormated := []entity.Purchase{
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
		}
		service.ClientService.(*mocks.IService).
			On("CreateArrayForInsert", mock.Anything).
			Return(&newPurchase, nil).
			Once()

		service.ClientService.(*mocks.IService).
			On("ValidationData", mock.Anything).
			Return(nil).
			Once()

		service.ClientService.(*mocks.IService).
			On("FormateDataForCreate", mock.Anything).
			Return(&purchasesFormated).
			Once()

		service.ClientService.(*mocks.IService).
			On("InsertBatch", mock.Anything, mock.Anything).
			Return(errors.New("Fail to insert")).
			Once()

		data, _ := os.ReadFile("../mocks/file/base_teste.txt")

		r := strings.NewReader(string(data))

		err := service.Handle(r, lengthBatch)
		assert.NotNil(t, err)

	})
}
