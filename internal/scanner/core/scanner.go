package core

import (
	log "github.com/ayushkr12/logger"
	"github.com/ayushkr12/spike/internal/pkg/db"
	"github.com/ayushkr12/spike/pkg/config"
)

type Scanner struct {
	Domains      []string
	SaveScanToDB bool
	Log          bool
	AppConfig    *config.AppConfig
	ToolsConfig  *config.ToolsConfig
}

func (s *Scanner) Scan() error {
	if err := s.ValidateScanner(); err != nil {
		return err
	}
	log.Debugf("Starting scan for domains: %v", s.Domains)
	if s.SaveScanToDB {
		db := &db.DB{}
		if err := db.Connect(s.AppConfig.DefaultDBPath); err != nil {
			return err
		}
		defer db.Close()
	}
	// todo
	return nil
}
