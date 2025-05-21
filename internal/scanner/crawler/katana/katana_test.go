package katana

import (
	"testing"

	"github.com/ayushkr12/spike/pkg/config"
)

func TestCrawlHosts(t *testing.T) {
	// Test with default values
	k := &config.KatanaConfig{
		CrawlDepth:         3,
		Threads:            10,
		ParallelismThreads: 10,
		Headless:           false,
		NoSandbox:          false,
		MaxCrawlTime:       "10s",
	}

	liveHosts := []string{
		"https://google.com",
	}

	urls, err := CrawlHosts(liveHosts, k)
	if err != nil {
		t.Fatalf("CrawlHosts() error = %v", err)
	}

	if len(urls) == 0 {
		t.Fatalf("CrawlHosts() returned no URLs")
	} else {
		t.Logf("CrawlHosts() returned %d URLs", len(urls))
	}
}
