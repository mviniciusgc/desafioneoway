package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatCpf(t *testing.T) {
	t.Run("success in format CPF", func(t *testing.T) {
		expected := "111.111.111-22"
		value := "11111111122"

		resp, err := FormatCpf(value)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, expected, *resp)

	})
	t.Run("success in format CPF with mask", func(t *testing.T) {
		expected := "111.111.111-22"
		value := "111.111.111-22"

		resp, err := FormatCpf(value)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, expected, *resp)

	})

}
func TestFormatCnpj(t *testing.T) {
	t.Run("success in format CNPJ", func(t *testing.T) {
		expected := "11.111.111/0001-22"
		value := "11111111000122"

		resp, err := FormatCnpj(value)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, expected, *resp)

	})
	t.Run("success in format CNPJ with mask", func(t *testing.T) {
		expected := "11.111.111/0001-22"
		value := "11.111.111/0001-22"

		resp, err := FormatCnpj(value)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, expected, *resp)

	})
}
