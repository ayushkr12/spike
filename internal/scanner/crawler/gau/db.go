package gau

import (
	"github.com/ayushkr12/spike/internal/pkg/db"
)

const SQL_CREATE_GAU_TABLE = `
CREATE TABLE IF NOT EXISTS gau (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	domain_id INTEGER NOT NULL,
	urls TEXT NOT NULL,
	FOREIGN KEY (domain_id) REFERENCES domains(id) ON DELETE CASCADE
);
`

const SQL_INSERT_GAU_URL = `INSERT INTO gau (domain_id, urls) VALUES (?, ?);`

func InsertURLs(db *db.DB, domainID int, urls []string) error {
	for _, url := range urls {
		if err := InsertURL(db, domainID, url); err != nil {
			return err
		}
	}
	return nil
}

func InsertURL(db *db.DB, domainID int, Url string) error {
	if err := db.ExecInsert(SQL_INSERT_GAU_URL, domainID, Url); err != nil {
		return err
	}
	return nil
}

func CreateGauTable(db *db.DB) error {
	if err := db.ExecStmt(SQL_CREATE_GAU_TABLE); err != nil {
		return err
	}
	return nil
}
