package utils

import (
	"errors"
	"strconv"
	"strings"
)

// Formamta os dados para Float64 caso o valor seja nulo retorna nil
func StringToFloat64(str string) (*float64, error) {
	if strings.ToLower(str) == "null" {
		return nil, errors.New("value is null")
	}
	str = strings.ReplaceAll(str, ",", ".")

	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

// Formata os dados para int caso o valor seja nulo retorna nil
func StringToInt(str string) (*int, error) {

	if str == "NULL" {
		return nil, errors.New("value is null")
	}
	value, err := strconv.Atoi(str)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func StringToBool(str string) (*bool, error) {

	if str == "NULL" {
		return nil, errors.New("value is null")
	}

	value, err := strconv.ParseBool(str)
	if err != nil {
		return nil, err
	}

	return &value, nil

}
