package tables

import (
	"database/sql"
)

// Criando a tabela e seus campos
func CreateTablePurchase(conn *sql.DB) error {
	sql := `CREATE TABLE IF NOT EXISTS purchase (id SERIAL PRIMARY KEY, cpf_cnpj VARCHAR(18) NOT NULL, private BOOLEAN, incomplete BOOLEAN, date_last_compra timestamp NULL, average_ticket FLOAT NULL, last_purchase_ticket FLOAT NULL, store_more_frequent VARCHAR(18) NULL, store_last_purchase VARCHAR(18) NULL)`

	_, err := conn.Exec(sql)
	if err != nil {
		return err
	}

	defer conn.Close()

	return nil
}
