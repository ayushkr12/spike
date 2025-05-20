package subfinder

import (
	"github.com/ayushkr12/spike/internal/pkg/db"
)

var SQL_CREATE_SUBDOMAINS_TABLE = `
CREATE TABLE IF NOT EXISTS subdomains (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	domain_id INTEGER NOT NULL,
	subdomain TEXT NOT NULL,
	FOREIGN KEY (domain_id) REFERENCES domains(id) ON DELETE CASCADE
);
`

var SQL_INSERT_SUBDOMAINS = `INSERT INTO subdomains (domain_id, subdomain) VALUES (?, ?);`

func CreateSubdomainsTable(db *db.DB) error {
	if err := db.ExecStmt(SQL_CREATE_SUBDOMAINS_TABLE); err != nil {
		return err
	}
	return nil
}

func InsertSubdomain(db *db.DB, domainID int, subdomain string) error {
	if err := db.ExecInsert(SQL_INSERT_SUBDOMAINS, domainID, subdomain); err != nil {
		return err
	}
	return nil
}

func InsertSubdomains(db *db.DB, domainID int, subdomains []string) error {
	for _, subdomain := range subdomains {
		if err := InsertSubdomain(db, domainID, subdomain); err != nil {
			return err
		}
	}
	return nil
}
