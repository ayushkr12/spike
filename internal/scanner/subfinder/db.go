package subfinder

import (
	"github.com/ayushkr12/spike/internal/pkg/db"
)

const SQL_CREATE_SUBDOMAINS_TABLE = `
CREATE TABLE IF NOT EXISTS subdomains (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	domain_id INTEGER NOT NULL,
	subdomain TEXT NOT NULL,
	FOREIGN KEY (domain_id) REFERENCES domains(id) ON DELETE CASCADE
);
`

const SQL_INSERT_SUBDOMAIN = `INSERT INTO subdomains (domain_id, subdomain) VALUES (?, ?);`

func InsertSubdomains(dbClient *db.DB, domainID int, subdomains []string) error {
	if len(subdomains) == 0 {
		return nil
	}
	// Prepare the SQL statement for bulk insert
	argsList := db.MakeArgsList(domainID, subdomains)

	// Execute the bulk insert
	if err := dbClient.ExecBulkInsert(SQL_INSERT_SUBDOMAIN, argsList); err != nil {
		return err
	}
	return nil
}

func InsertSubdomain(db *db.DB, domainID int, subdomain string) error {
	if err := db.ExecInsert(SQL_INSERT_SUBDOMAIN, domainID, subdomain); err != nil {
		return err
	}
	return nil
}

func CreateSubdomainsTable(db *db.DB) error {
	if err := db.ExecStmt(SQL_CREATE_SUBDOMAINS_TABLE); err != nil {
		return err
	}
	return nil
}
