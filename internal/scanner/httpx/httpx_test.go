package httpx

import (
	"testing"
)

func TestProbe(t *testing.T) {
	// Test with default values
	hosts := []string{"https://example.com"}
	threads := 50

	urls, err := Probe(hosts, threads)
	if err != nil {
		t.Fatalf("Probe() error = %v", err)
	}

	if len(urls) == 0 {
		t.Fatalf("Probe() returned no URLs")
	} else {
		t.Logf("Probe() returned %s URLs", urls)
	}
}
