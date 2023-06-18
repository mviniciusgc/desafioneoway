package repositorie

import "github.com/mviniciusgc/desafio_neoway/src/entity"

type IRepository interface {
	Create(lineFile []entity.Purchase) (id *int64, err error)
}
