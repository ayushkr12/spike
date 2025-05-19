package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

func (db *DB) ExecStmt(stmt string) error {
	_, err := db.Exec(stmt)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) Connect(dbPath string) error {
	var err error
	db.DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) Close() error {
	if err := db.DB.Close(); err != nil {
		return err
	}
	return nil
}
