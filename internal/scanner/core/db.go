package core

import (
	"fmt"

	"github.com/ayushkr12/spike/internal/pkg/db"
)

const SQL_CREATE_DOMAINS_TABLE = `
CREATE TABLE IF NOT EXISTS domains (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	domain TEXT NOT NULL UNIQUE,
	is_scanned BOOLEAN DEFAULT FALSE
);
`

const SQL_INSERT_DOMAIN = `INSERT INTO domains (domain) VALUES (?)`
const SQL_UPDATE_DOMAIN_SCANNED = `UPDATE domains SET is_scanned = TRUE WHERE domain = ?`
const SQL_SELECT_DOMAINS = `SELECT id, domain, is_scanned FROM domains`

type DomainConfig struct {
	Id        int
	Domain    string
	IsScanned bool
}

func SelectDomains(db *db.DB) ([]DomainConfig, error) {
	rows, err := db.Query(SQL_SELECT_DOMAINS)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var domains []DomainConfig
	for rows.Next() {
		var domain DomainConfig
		if err := rows.Scan(&domain.Id, &domain.Domain, &domain.IsScanned); err != nil {
			return nil, err
		}
		domains = append(domains, domain)
	}
	return domains, nil
}

func UpdateDomainScanned(db *db.DB, domainId int) error {
	if err := db.ExecInsert(SQL_UPDATE_DOMAIN_SCANNED, domainId); err != nil {
		return err
	}
	return nil
}

func PushDomainsToDB(dbClient *db.DB, domains []string) error {
	if err := dbClient.ExecStmt(SQL_CREATE_DOMAINS_TABLE); err != nil {
		return fmt.Errorf("failed to create domains table: %w", err)
	}
	if err := InsertDomains(dbClient, domains); err != nil {
		return fmt.Errorf("failed to insert domains: %w", err)
	}
	return nil
}

func InsertDomains(dbClient *db.DB, domains []string) error {
	argsList := make([][]any, len(domains))
	for i, domain := range domains {
		argsList[i] = []any{domain}
	}
	if err := dbClient.ExecBulkInsert(SQL_INSERT_DOMAIN, argsList); err != nil {
		return err
	}
	return nil
}

func InsertDomain(db *db.DB, domain string) error {
	if err := db.ExecInsert(SQL_INSERT_DOMAIN, domain); err != nil {
		return err
	}
	return nil
}
