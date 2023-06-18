package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidadeCnpj(t *testing.T) {
	t.Run("success when CNPJ is valid", func(t *testing.T) {
		value := "23.359.106/0001-81"

		resp := ValidadeCnpj(value)
		assert.Equal(t, true, resp)

	})
	t.Run("success when CNPJ is valid and not mask", func(t *testing.T) {
		value := "23359106000181"

		resp := ValidadeCnpj(value)
		assert.Equal(t, true, resp)

	})
	t.Run("error when CNPJ is valid", func(t *testing.T) {
		value := "11.111.111/0001-22"

		resp := ValidadeCnpj(value)
		assert.Equal(t, false, resp)

	})
}

func TestValidadeCpf(t *testing.T) {
	t.Run("success when CPF is valid", func(t *testing.T) {
		value := "209.516.480-63"

		resp := ValidadeCpf(value)
		assert.Equal(t, true, resp)

	})
	t.Run("error when CPF is valid", func(t *testing.T) {
		value := "111.111.111-111"

		resp := ValidadeCpf(value)
		assert.Equal(t, false, resp)

	})
}
