package httpx

import (
	"github.com/ayushkr12/spike/internal/pkg/db"
)

var SQL_CREATE_LIVE_HOSTS_TABLE = `
CREATE TABLE IF NOT EXISTS live_hosts (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	domain_id INTEGER NOT NULL,
	live_host TEXT NOT NULL,
	FOREIGN KEY (domain_id) REFERENCES domains(id) ON DELETE CASCADE
);
`

var SQL_INSERT_LIVE_HOSTS = `INSERT INTO live_hosts (domain_id, live_host) VALUES (?, ?);`

func CreateLiveHostsTable(db *db.DB) error {
	if err := db.ExecStmt(SQL_CREATE_LIVE_HOSTS_TABLE); err != nil {
		return err
	}
	return nil
}

func InsertLiveHost(db *db.DB, domainID int, liveHost string) error {
	if err := db.ExecInsert(SQL_INSERT_LIVE_HOSTS, domainID, liveHost); err != nil {
		return err
	}
	return nil
}
