package service

import (
	"strings"

	"github.com/mviniciusgc/desafio_neoway/src/entity"
	"github.com/mviniciusgc/desafio_neoway/src/utils"
)

// Formata os dados para serem inseridos no banco
func (c *ClientService) FormateDataForCreate(values []string) *[]entity.Purchase {

	// Cria um array de struct para receber os dados
	purchases := []entity.Purchase{}

	for _, value := range values {
		purchase := entity.Purchase{}

		line := strings.Fields(value)
		//formata cpf ou cnpj
		cpfCpnj, _ := formateCpfCnpj(line[0])

		purchase.CPFCNPJ = cpfCpnj

		//convertendo string para bool
		private, _ := utils.StringToBool(line[1])

		purchase.Private = private

		//convertendo string para bool
		incomplete, _ := utils.StringToBool(line[2])

		purchase.Incomplete = incomplete

		if line[3] != "NULL" {
			purchase.DateLastCompra = &line[3]
		}

		//convertendo string para float
		averageTicket, _ := utils.StringToFloat64(line[4])

		purchase.AverageTicket = averageTicket

		//convertendo string para float
		lastPurchaseTicket, _ := utils.StringToFloat64(line[5])

		purchase.LastPurchaseTicket = lastPurchaseTicket

		// Formata os dados para serem inseridos no banco
		cpfCpnj, _ = formateCpfCnpj(line[6])

		purchase.StoreMoreFrequent = cpfCpnj

		// Formata os dados para serem inseridos no banco
		cpfCpnj, _ = formateCpfCnpj(line[7])

		purchase.StoreLastPurchase = cpfCpnj

		purchases = append(purchases, purchase)
	}

	return &purchases

}

func formateCpfCnpj(value string) (*string, error) {

	// rever a parte de validação de cpf e cnpj
	if len(value) == 11 {
		resp, err := utils.FormatCpf(value)
		if err != nil {
			return nil, err
		}
		return resp, nil

	}

	resp, err := utils.FormatCnpj(value)
	if err != nil {
		return nil, err
	}

	return resp, nil

}
