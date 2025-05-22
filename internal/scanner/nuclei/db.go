package nuclei

import (
	"github.com/ayushkr12/spike/internal/pkg/db"
)

const SQL_CREATE_NUCLEI_TABLE = `
CREATE TABLE IF NOT EXISTS nuclei (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	domain_id INTEGER NOT NULL,
	report TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (domain_id) REFERENCES domains(id) ON DELETE CASCADE
);
`

const SQL_INSERT_NUCLEI_REPORT = `INSERT INTO nuclei (domain_id, report) VALUES (?, ?);`

func InsertReports(dbClient *db.DB, domainID int, reports []string) error {
	if len(reports) == 0 {
		return nil
	}
	// Prepare the SQL statement for bulk insert
	argsList := db.MakeArgsList(domainID, reports)

	// Execute the bulk insert
	if err := dbClient.ExecBulkInsert(SQL_INSERT_NUCLEI_REPORT, argsList); err != nil {
		return err
	}
	return nil
}

func InsertReport(db *db.DB, domainID int, report string) error {
	if err := db.ExecInsert(SQL_INSERT_NUCLEI_REPORT, domainID, report); err != nil {
		return err
	}
	return nil
}

func CreateNucleiTable(db *db.DB) error {
	if err := db.ExecStmt(SQL_CREATE_NUCLEI_TABLE); err != nil {
		return err
	}
	return nil
}
