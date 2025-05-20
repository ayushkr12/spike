package katana

import (
	"github.com/ayushkr12/spike/internal/pkg/db"
)

var SQL_CREATE_KATANA_TABLE = `
CREATE TABLE IF NOT EXISTS katana (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	domain_id INTEGER NOT NULL,
	urls TEXT NOT NULL,
	FOREIGN KEY (domain_id) REFERENCES domains(id) ON DELETE CASCADE
);
`

var SQL_INSERT_KATANA_URL = `INSERT INTO katana (domain_id, urls) VALUES (?, ?);`

func InsertURLs(db *db.DB, domainID int, urls []string) error {
	for _, url := range urls {
		if err := InsertURL(db, domainID, url); err != nil {
			return err
		}
	}
	return nil
}

func InsertURL(db *db.DB, domainID int, Url string) error {
	if err := db.ExecInsert(SQL_INSERT_KATANA_URL, domainID, Url); err != nil {
		return err
	}
	return nil
}

func CreateKatanaTable(db *db.DB) error {
	if err := db.ExecStmt(SQL_CREATE_KATANA_TABLE); err != nil {
		return err
	}
	return nil
}
