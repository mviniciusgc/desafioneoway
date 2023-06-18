package controller

import (
	"io"
)

// responsavel por orquestar as chamadas de funcoes
func (c *ClientController) Handle(purchase io.Reader, lengthBatch int) error {

	newPurchase, err := c.ClientService.CreateArrayForInsert(purchase)
	if err != nil {
		return err
	}

	err = c.ClientService.ValidationData(*newPurchase)
	if err != nil {
		return err
	}

	purchasesFormated := c.ClientService.FormateDataForCreate(*newPurchase)

	err = c.ClientService.InsertBatch(purchasesFormated, lengthBatch)
	if err != nil {
		return err
	}

	return nil
}
