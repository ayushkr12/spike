package core

import (
	"fmt"

	log "github.com/ayushkr12/logger"
)

func (s *Scanner) ValidateScanner() error {
	if s.Domains == nil {
		return fmt.Errorf("domains cannot be nil")
	}
	if !s.SaveScanToDB && !s.Log {
		log.Warn("Neither SaveScanToDB nor Log is enabled. Nothing to do, enabling Log as a fallback.")
		s.Log = true
	}
	return nil
}
