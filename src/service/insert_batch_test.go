package service

import (
	"testing"

	"github.com/mviniciusgc/desafio_neoway/src/entity"
	"github.com/mviniciusgc/desafio_neoway/src/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupController() *ClientService {
	as := InitializeService(&mocks.IRepository{})

	return as
}
func TestHandle(t *testing.T) {

	t.Run("success to insert in batch", func(t *testing.T) {
		service := setupController()
		lengthBatch := 5

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
		service.clientrepository.(*mocks.IRepository).
			On("Create", mock.Anything).
			Return(nil, nil).
			Once()

		err := service.InsertBatch(&purchasesFormated, lengthBatch)
		assert.Nil(t, err)

	})

	t.Run("success to insert in batch more data", func(t *testing.T) {
		service := setupController()
		lengthBatch := 5

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
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
		}
		service.clientrepository.(*mocks.IRepository).
			On("Create", mock.Anything).
			Return(nil, nil)

		err := service.InsertBatch(&purchasesFormated, lengthBatch)
		assert.Nil(t, err)

	})

}
