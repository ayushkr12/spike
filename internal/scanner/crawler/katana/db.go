package katana

import (
	"github.com/ayushkr12/spike/internal/pkg/db"
)

const SQL_CREATE_KATANA_TABLE = `
CREATE TABLE IF NOT EXISTS katana (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	domain_id INTEGER NOT NULL,
	url TEXT NOT NULL,
	FOREIGN KEY (domain_id) REFERENCES domains(id) ON DELETE CASCADE
);
`

const SQL_INSERT_KATANA_URL = `INSERT INTO katana (domain_id, url) VALUES (?, ?);`

func InsertURLs(dbClient *db.DB, domainID int, urls []string) error {
	if len(urls) == 0 {
		return nil
	}
	// Prepare the SQL statement for bulk insert
	argsList := db.MakeArgsList(domainID, urls)

	// Execute the bulk insert
	if err := dbClient.ExecBulkInsert(SQL_INSERT_KATANA_URL, argsList); err != nil {
		return err
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
