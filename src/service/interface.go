package service

import (
	"io"

	"github.com/mviniciusgc/desafio_neoway/src/entity"
)

type IService interface {
	CreateArrayForInsert(purchase io.Reader) (*[]string, error)
	FormateDataForCreate(values []string) *[]entity.Purchase
	ValidationData(values []string) error
	InsertBatch(Purchases *[]entity.Purchase, lengthBatch int) error
}
