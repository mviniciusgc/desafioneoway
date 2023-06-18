package repositorie

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/mviniciusgc/desafio_neoway/src/entity"
	"github.com/mviniciusgc/desafio_neoway/src/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRepository() *ClientRepository {
	as := InitializeClienteRepository(&mocks.IClientDB{})

	return as
}

// TODO: continuar
func TestCreate(t *testing.T) {
	t.Run("success in formatting data to save in db", func(t *testing.T) {
		service := setupRepository()

		CPFCNPJ := "123.456.789-00"
		Private := bool(true)
		Incomplete := bool(false)
		DateLastCompra := "2016-01-02"
		AverageTicket := float64(123.45)
		LastPurchaseTicket := float64(123.45)
		StoreMoreFrequent := "79.379.491/0001-83"
		StoreLastPurchase := "79.379.491/0001-83"

		lineFile := []entity.Purchase{
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
		}

		sqlDB := sql.DB{}
		sqlTx := sql.Tx{}
		sqlStmt := sql.Stmt{}
		var sqlResult sql.Result

		service.db.(*mocks.IClientDB).
			On("InitDB").
			Return(&sqlDB, nil).
			Once()

		service.db.(*mocks.IClientDB).
			On("Begin", mock.Anything, mock.Anything).
			Return(&sqlTx, nil).
			Once()

		service.db.(*mocks.IClientDB).
			On("Prepare", mock.Anything, mock.Anything, mock.Anything).
			Return(&sqlStmt, nil).
			Once()

		service.db.(*mocks.IClientDB).
			On("Exec", mock.Anything, mock.Anything).
			Return(sqlResult, nil)

		service.db.(*mocks.IClientDB).
			On("Close", mock.Anything).
			Return(nil).
			Once()

		service.db.(*mocks.IClientDB).
			On("Commit", mock.Anything).
			Return(nil).
			Once()

		resp, err := service.Create(lineFile)
		assert.Nil(t, err)
		assert.Nil(t, resp)

	})

	t.Run("error to initDB", func(t *testing.T) {
		service := setupRepository()

		CPFCNPJ := "123.456.789-00"
		Private := bool(true)
		Incomplete := bool(false)
		DateLastCompra := "2016-01-02"
		AverageTicket := float64(123.45)
		LastPurchaseTicket := float64(123.45)
		StoreMoreFrequent := "79.379.491/0001-83"
		StoreLastPurchase := "79.379.491/0001-83"

		lineFile := []entity.Purchase{
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
		}

		service.db.(*mocks.IClientDB).
			On("InitDB").
			Return(nil, errors.New("fail to init db")).
			Once()

		resp, err := service.Create(lineFile)
		assert.NotNil(t, err)
		assert.Nil(t, resp)

	})

	t.Run("error when initial transation", func(t *testing.T) {
		service := setupRepository()

		CPFCNPJ := "123.456.789-00"
		Private := bool(true)
		Incomplete := bool(false)
		DateLastCompra := "2016-01-02"
		AverageTicket := float64(123.45)
		LastPurchaseTicket := float64(123.45)
		StoreMoreFrequent := "79.379.491/0001-83"
		StoreLastPurchase := "79.379.491/0001-83"

		lineFile := []entity.Purchase{
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
		}

		sqlDB := sql.DB{}

		service.db.(*mocks.IClientDB).
			On("InitDB").
			Return(&sqlDB, nil).
			Once()

		service.db.(*mocks.IClientDB).
			On("Begin", mock.Anything, mock.Anything).
			Return(nil, errors.New("error to init transation")).
			Once()

		resp, err := service.Create(lineFile)
		assert.NotNil(t, err)
		assert.Nil(t, resp)

	})

	t.Run("error to Prepare ", func(t *testing.T) {
		service := setupRepository()

		CPFCNPJ := "123.456.789-00"
		Private := bool(true)
		Incomplete := bool(false)
		DateLastCompra := "2016-01-02"
		AverageTicket := float64(123.45)
		LastPurchaseTicket := float64(123.45)
		StoreMoreFrequent := "79.379.491/0001-83"
		StoreLastPurchase := "79.379.491/0001-83"

		lineFile := []entity.Purchase{
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
		}

		sqlDB := sql.DB{}
		sqlTx := sql.Tx{}

		service.db.(*mocks.IClientDB).
			On("InitDB").
			Return(&sqlDB, nil).
			Once()

		service.db.(*mocks.IClientDB).
			On("Begin", mock.Anything, mock.Anything).
			Return(&sqlTx, nil).
			Once()

		service.db.(*mocks.IClientDB).
			On("Prepare", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("error to prepare")).
			Once()

		resp, err := service.Create(lineFile)
		assert.NotNil(t, err)
		assert.Nil(t, resp)

	})

	t.Run("error to execute", func(t *testing.T) {
		service := setupRepository()

		CPFCNPJ := "123.456.789-00"
		Private := bool(true)
		Incomplete := bool(false)
		DateLastCompra := "2016-01-02"
		AverageTicket := float64(123.45)
		LastPurchaseTicket := float64(123.45)
		StoreMoreFrequent := "79.379.491/0001-83"
		StoreLastPurchase := "79.379.491/0001-83"

		lineFile := []entity.Purchase{
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
		}

		sqlDB := sql.DB{}
		sqlTx := sql.Tx{}
		sqlStmt := sql.Stmt{}

		service.db.(*mocks.IClientDB).
			On("InitDB").
			Return(&sqlDB, nil).
			Once()

		service.db.(*mocks.IClientDB).
			On("Begin", mock.Anything, mock.Anything).
			Return(&sqlTx, nil).
			Once()

		service.db.(*mocks.IClientDB).
			On("Prepare", mock.Anything, mock.Anything, mock.Anything).
			Return(&sqlStmt, nil).
			Once()

		service.db.(*mocks.IClientDB).
			On("Exec", mock.Anything, mock.Anything).
			Return(nil, errors.New("error to execute"))

		resp, err := service.Create(lineFile)
		assert.NotNil(t, err)
		assert.Nil(t, resp)

	})

	t.Run("error to close transaction", func(t *testing.T) {
		service := setupRepository()

		CPFCNPJ := "123.456.789-00"
		Private := bool(true)
		Incomplete := bool(false)
		DateLastCompra := "2016-01-02"
		AverageTicket := float64(123.45)
		LastPurchaseTicket := float64(123.45)
		StoreMoreFrequent := "79.379.491/0001-83"
		StoreLastPurchase := "79.379.491/0001-83"

		lineFile := []entity.Purchase{
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
		}

		sqlDB := sql.DB{}
		sqlTx := sql.Tx{}
		sqlStmt := sql.Stmt{}
		var sqlResult sql.Result

		service.db.(*mocks.IClientDB).
			On("InitDB").
			Return(&sqlDB, nil).
			Once()

		service.db.(*mocks.IClientDB).
			On("Begin", mock.Anything, mock.Anything).
			Return(&sqlTx, nil).
			Once()

		service.db.(*mocks.IClientDB).
			On("Prepare", mock.Anything, mock.Anything, mock.Anything).
			Return(&sqlStmt, nil).
			Once()

		service.db.(*mocks.IClientDB).
			On("Exec", mock.Anything, mock.Anything).
			Return(sqlResult, nil)

		service.db.(*mocks.IClientDB).
			On("Close", mock.Anything).
			Return(errors.New("error to close")).
			Once()

		resp, err := service.Create(lineFile)
		assert.NotNil(t, err)
		assert.Nil(t, resp)

	})

	t.Run("error to commit", func(t *testing.T) {
		service := setupRepository()

		CPFCNPJ := "123.456.789-00"
		Private := bool(true)
		Incomplete := bool(false)
		DateLastCompra := "2016-01-02"
		AverageTicket := float64(123.45)
		LastPurchaseTicket := float64(123.45)
		StoreMoreFrequent := "79.379.491/0001-83"
		StoreLastPurchase := "79.379.491/0001-83"

		lineFile := []entity.Purchase{
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
			{CPFCNPJ: &CPFCNPJ, Private: &Private, Incomplete: &Incomplete, DateLastCompra: &DateLastCompra, AverageTicket: &AverageTicket, LastPurchaseTicket: &LastPurchaseTicket, StoreMoreFrequent: &StoreMoreFrequent, StoreLastPurchase: &StoreLastPurchase},
		}

		sqlDB := sql.DB{}
		sqlTx := sql.Tx{}
		sqlStmt := sql.Stmt{}
		var sqlResult sql.Result

		service.db.(*mocks.IClientDB).
			On("InitDB").
			Return(&sqlDB, nil).
			Once()

		service.db.(*mocks.IClientDB).
			On("Begin", mock.Anything, mock.Anything).
			Return(&sqlTx, nil).
			Once()

		service.db.(*mocks.IClientDB).
			On("Prepare", mock.Anything, mock.Anything, mock.Anything).
			Return(&sqlStmt, nil).
			Once()

		service.db.(*mocks.IClientDB).
			On("Exec", mock.Anything, mock.Anything).
			Return(sqlResult, nil)

		service.db.(*mocks.IClientDB).
			On("Close", mock.Anything).
			Return(nil).
			Once()

		service.db.(*mocks.IClientDB).
			On("Commit", mock.Anything).
			Return(errors.New("error to commit")).
			Once()

		resp, err := service.Create(lineFile)
		assert.NotNil(t, err)
		assert.Nil(t, resp)

	})

}
