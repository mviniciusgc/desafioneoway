package service

import (
	"errors"
	"strings"

	"github.com/mviniciusgc/desafio_neoway/src/utils"
)

func validateCpfCnpj(value string) (bool, error) {
	lenthCpf := 11

	if len(value) == lenthCpf {
		cpfValid := utils.ValidadeCpf(value)
		if !cpfValid {
			return false, errors.New("CPF invalid")
		}

	} else if len(value) == 14 {

		cnpjValid := utils.ValidadeCnpj(value)
		if !cnpjValid {
			return false, errors.New("CNPJ invalid")
		}
	}

	return true, nil

}

// Formata os dados para serem inseridos no banco
func (c *ClientService) ValidationData(values []string) error {

	// validando se os valores estão corretos
	for _, value := range values {

		line := strings.Fields(value)

		//validando cpf ou cnpj
		_, err := validateCpfCnpj(line[0])
		if err != nil {
			return err
		}

		//validando se o valor é bool
		_, err = utils.StringToBool(line[1])
		if err != nil {
			return err
		}

		//validando se o valor é bool
		_, err = utils.StringToBool(line[2])
		if err != nil {
			return err
		}

		//validando se o valor é um float
		_, err = utils.StringToFloat64(line[4])
		if err != nil && err.Error() != "value is null" {
			return err
		}

		//validando se o valor é um float
		_, err = utils.StringToFloat64(line[5])
		if err != nil && err.Error() != "value is null" {
			return err
		}

		//validando cpf ou cnpj
		_, err = validateCpfCnpj(line[6])
		if err != nil {
			return err
		}

		// //validando cpf ou cnpj
		_, err = validateCpfCnpj(line[7])
		if err != nil {
			return err
		}

	}

	return nil

}
