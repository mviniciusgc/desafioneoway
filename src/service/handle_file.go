package service

import (
	"bufio"
	"io"
	"io/ioutil"
	"strings"
)

func (c *ClientService) CreateArrayForInsert(purchase io.Reader) (*[]string, error) {
	var lines []string

	// Lê o arquivo recebido
	date, err := ioutil.ReadAll(purchase)
	if err != nil {
		return nil, err
	}

	content := string(date)

	content = strings.ReplaceAll(content, ".", "")
	content = strings.ReplaceAll(content, "-", "")
	content = strings.ReplaceAll(content, "/", "")

	// Cria um scanner para ler o arquivo linha a linha
	scanner := bufio.NewScanner(strings.NewReader(content))

	// Ignorar a primeira linha do arquivo já que é o cabeçalho
	scanner.Scan()

	// Lê o arquivo linha a linha facilita a manipulação dos dados
	for scanner.Scan() {

		lineStr := scanner.Text()

		lines = append(lines, lineStr)

	}

	return &lines, nil
}
