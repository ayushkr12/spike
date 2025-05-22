package core

import (
	"fmt"

	log "github.com/ayushkr12/logger"
	"github.com/ayushkr12/spike/internal/scanner/crawler/gau"
	"github.com/ayushkr12/spike/internal/scanner/crawler/katana"
	"github.com/ayushkr12/spike/pkg/config"
)

func RunCrawlers(c *config.CrawlerConfig, liveHosts []string, rootDomain string) ([]string, error) {
	if err := validateCrawlerConfig(c); err != nil {
		return nil, fmt.Errorf("invalid crawler config: %v", err)
	}

	var crawledUrls []string
	if c.Katana.Enabled {
		log.Infof("Running katana on %d live hosts", len(liveHosts))
		urls, err := katana.CrawlHosts(liveHosts, &c.Katana)
		if err != nil {
			return nil, fmt.Errorf("failed to run katana: %v", err)
		}
		crawledUrls = append(crawledUrls, urls...)
	}

	if c.Gau.Enabled {
		log.Infof("Running gau on root domain: %s", rootDomain)
		urls, err := gau.GetAllUrls(rootDomain, c.Gau.Threads)
		if err != nil {
			return nil, fmt.Errorf("failed to run gau: %v", err)
		}
		crawledUrls = append(crawledUrls, urls...)
	}

	return crawledUrls, nil
}
