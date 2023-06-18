package db

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/mviniciusgc/desafio_neoway/src/db/tables"
	"github.com/spf13/viper"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type IDataBase interface {
	InitDB() (*sql.DB, error)
}

func (c *ClientDB) InitDB() (*sql.DB, error) {
	credentials := getCredentialConnection()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", credentials.Host, credentials.Port, credentials.User, credentials.Password, credentials.DBName)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, err
}

// responsavel por criar as tabelas
func (c *ClientDB) InitializeTables() error {
	conn, err := c.InitDB()
	if err != nil {
		return err
	}

	err = tables.CreateTablePurchase(conn)
	if err != nil {
		return err
	}

	return nil
}

func (c *ClientDB) Begin(conn *sql.DB) (*sql.Tx, error) {
	tx, err := conn.Begin()
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *ClientDB) Prepare(tx *sql.Tx, nameTable string, fields []string) (*sql.Stmt, error) {

	stmt, err := tx.Prepare(pq.CopyIn(nameTable, fields...))
	if err != nil {
		return nil, err
	}

	return stmt, nil
}

func (c *ClientDB) Close(smtp *sql.Stmt) error {
	err := smtp.Close()
	if err != nil {
		return err
	}

	return nil
}

func (c *ClientDB) Commit(tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (c *ClientDB) Exec(stmt *sql.Stmt, lineFile []interface{}) (sql.Result, error) {
	resp, err := stmt.Exec(lineFile...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func getCredentialConnection() DBConfig {

	return DBConfig{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
		DBName:   viper.GetString("DB_NAME"),
	}

}
