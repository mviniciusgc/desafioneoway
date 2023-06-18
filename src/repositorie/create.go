package repositorie

import (
	"github.com/mviniciusgc/desafio_neoway/src/entity"
)

func (c *ClientRepository) Create(lineFile []entity.Purchase) (id *int64, err error) {

	nameTable := "purchase"
	fields := []string{"cpf_cnpj", "private", "incomplete", "date_last_compra", "average_ticket", "last_purchase_ticket", "store_more_frequent", "store_last_purchase"}

	//Iniciando a conexão com o banco
	conn, err := c.db.InitDB()
	if err != nil {
		return nil, err
	}

	//iniciano a transação
	tx, err := c.db.Begin(conn)
	if err != nil {
		return nil, err
	}

	// Cria uma instrução dentro de uma transação para ser executada
	stmt, err := c.db.Prepare(tx, nameTable, fields)
	if err != nil {
		return nil, err
	}

	// vai executar a instrução para cada linha do arquivo
	for _, line := range lineFile {
		_, err = c.db.Exec(stmt, []interface{}{line.CPFCNPJ, line.Private, line.Incomplete, line.DateLastCompra, line.AverageTicket, line.LastPurchaseTicket, line.StoreMoreFrequent, line.StoreLastPurchase})
		if err != nil {
			return nil, err
		}
	}

	// Fechando a instrução
	err = c.db.Close(stmt)
	if err != nil {
		return nil, err
	}

	// confirmando a transação
	err = c.db.Commit(tx)
	if err != nil {
		return nil, err
	}

	return nil, err
}
