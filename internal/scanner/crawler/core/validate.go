package crawler

import (
	"fmt"

	"github.com/ayushkr12/spike/pkg/config"
)

func validateCrawlerConfig(c *config.CrawlerConfig) error {
	if !c.Katana.Enabled && !c.Gau.Enabled {
		return fmt.Errorf("at least one crawler must be enabled")
	}

	return nil
}
