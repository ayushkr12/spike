package katana

import (
	"testing"
)

func TestCrawlHosts(t *testing.T) {
	// Test with default values
	k := &KatanaScanner{
		LiveHosts:         []string{"https://example.com"},
		CrawlDepth:        3,
		Threads:           10,
		ParllelismThreads: 10,
		Headless:          false,
		NoSandbox:         false,
		MaxCrawlTime:      "10s",
	}

	urls, err := CrawlHosts(k)
	if err != nil {
		t.Fatalf("CrawlHosts() error = %v", err)
	}

	if len(urls) == 0 {
		t.Fatalf("CrawlHosts() returned no URLs")
	} else {
		t.Logf("CrawlHosts() returned %d URLs", len(urls))
	}
}
