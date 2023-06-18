package db

import "database/sql"

type IClientDB interface {
	InitializeTables() error
	InitDB() (*sql.DB, error)
	Begin(conn *sql.DB) (*sql.Tx, error)
	Prepare(tx *sql.Tx, nameTable string, fields []string) (*sql.Stmt, error)
	Close(smtp *sql.Stmt) error
	Commit(tx *sql.Tx) error
	Exec(stmt *sql.Stmt, lineFile []interface{}) (sql.Result, error)
}
