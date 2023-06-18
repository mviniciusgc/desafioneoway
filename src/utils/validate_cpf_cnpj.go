package utils

import (
	"github.com/klassmann/cpfcnpj"
)

func ValidadeCnpj(cnpj string) bool {

	resp := cpfcnpj.NewCNPJ(cnpj)

	return resp.IsValid()

}

func ValidadeCpf(cpf string) bool {
	resp := cpfcnpj.NewCPF(cpf)

	return resp.IsValid()

}
