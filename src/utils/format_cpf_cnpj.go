package utils

import (
	"regexp"

	"github.com/klassmann/cpfcnpj"
)

// Formata o cpf colocando mascara apartir do regex
func FormatCpf(cpf string) (*string, error) {

	expr, err := regexp.Compile(cpfcnpj.CPFFormatPattern)
	if err != nil {
		return nil, err
	}

	resp := expr.ReplaceAllString(cpf, "$1.$2.$3-$4")

	return &resp, nil
}

// Formata o cnpj colocando mascara apartir do regex
func FormatCnpj(cnpj string) (*string, error) {

	expr, err := regexp.Compile(cpfcnpj.CNPJFormatPattern)
	if err != nil {
		return nil, err
	}

	resp := expr.ReplaceAllString(cnpj, "$1.$2.$3/$4-$5")

	return &resp, nil
}
