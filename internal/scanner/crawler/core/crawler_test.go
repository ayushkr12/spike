package crawler

import (
	"testing"

	"github.com/ayushkr12/spike/pkg/config"
)

func TestRunCrawlers(t *testing.T) {
	c := &config.CrawlerConfig{
		Gau: config.GauConfig{
			Enabled: true,
			Threads: 10,
		},
		Katana: config.KatanaConfig{
			Enabled:            true,
			CrawlDepth:         3,
			Threads:            10,
			ParallelismThreads: 10,
			Headless:           false,
			NoSandbox:          false,
			MaxCrawlTime:       "10s",
		},
	}

	liveHosts := []string{"https://example.org"}
	rootDomain := "example.org"

	urls, err := RunCrawlers(c, liveHosts, rootDomain)
	if err != nil {
		t.Fatalf("RunCrawlers() error = %v", err)
	}

	if len(urls) == 0 {
		t.Fatalf("RunCrawlers() returned no URLs")
	} else {
		t.Logf("RunCrawlers() returned %d URLs", len(urls))
	}
}
