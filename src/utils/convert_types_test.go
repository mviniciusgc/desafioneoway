package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertStringToFloat64(t *testing.T) {
	t.Run("success in convert string for Float64", func(t *testing.T) {
		value := "1,1"

		resp, err := StringToFloat64(value)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, 1.1, *resp)

	})
	t.Run("success in convert string for Float64", func(t *testing.T) {
		value := "2.1"

		resp, err := StringToFloat64(value)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, 2.1, *resp)

	})
	t.Run("return error in convert invalid string for Float64", func(t *testing.T) {
		value := "a.1"

		resp, err := StringToFloat64(value)
		assert.NotNil(t, err)
		assert.Nil(t, resp)

	})
	t.Run("return error in convert value null in string for Float64", func(t *testing.T) {
		value := "NULL"

		resp, err := StringToFloat64(value)
		assert.NotNil(t, err)
		assert.Nil(t, resp)

	})
}
func TestConvertStringToInt(t *testing.T) {
	t.Run("success in convert string for int", func(t *testing.T) {
		value := "1"

		resp, err := StringToInt(value)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, 1, *resp)

	})
	t.Run("return error in convert invalid string for int", func(t *testing.T) {
		value := "a.1"

		resp, err := StringToInt(value)
		assert.NotNil(t, err)
		assert.Nil(t, resp)

	})
	t.Run("return eror in convert value null in string for int", func(t *testing.T) {
		value := "NULL"

		resp, err := StringToInt(value)
		assert.NotNil(t, err)
		assert.Nil(t, resp)

	})

}
func TestConvertStringToBool(t *testing.T) {
	t.Run("success in convert string for bool 'True' ", func(t *testing.T) {
		value := "1"

		resp, err := StringToBool(value)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, true, *resp)

	})
	t.Run("success in convert string for bool 'False' ", func(t *testing.T) {
		value := "0"

		resp, err := StringToBool(value)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, false, *resp)

	})
	t.Run("return error in convert invalid string for bool", func(t *testing.T) {
		value := "a.1"

		resp, err := StringToBool(value)
		assert.NotNil(t, err)
		assert.Nil(t, resp)

	})
	t.Run("return error in convert value null in string for bool", func(t *testing.T) {
		value := "NULL"

		resp, err := StringToBool(value)
		assert.NotNil(t, err)
		assert.Nil(t, resp)

	})
}
