package service

import (
	"math"

	"github.com/mviniciusgc/desafio_neoway/src/entity"
)

func (c *ClientService) InsertBatch(Purchases *[]entity.Purchase, lengthBatch int) error {

	insertPurchase := *Purchases
	lengthPurchase := len(*Purchases)
	// Calcula quantos lotes serão inseridos no banco arredondando para cima pois pode dar um valor quebrado
	result := float64(lengthPurchase) / float64(lengthBatch)
	roundedResult := math.Ceil(result)

	// Inicia a inserção no banco
	for i := 0; i < int(roundedResult); i++ {
		smallArrayLines := make([]entity.Purchase, lengthBatch)

		// Se o tamanho do array for menor que o tamanho do lote ele insere o que tem no array seria o ultimo lote
		if lengthPurchase < lengthBatch {
			_, err := c.clientrepository.Create(insertPurchase)
			if err != nil {
				return err
			}

			break
		}

		// Copia o array para um array menor para ser inserido no banco
		copy(smallArrayLines, insertPurchase[0:lengthBatch])
		insertPurchase = append(insertPurchase[:0], insertPurchase[lengthBatch:]...)

		// chama o repositorio para inserir no banco
		_, err := c.clientrepository.Create(smallArrayLines)
		if err != nil {
			return err
		}
		lengthPurchase = len(insertPurchase)
	}

	return nil
}
