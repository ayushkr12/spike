package katana

import (
	"strconv"

	"github.com/ayushkr12/spike/internal/pkg/utils"
)

type KatanaScanner struct {
	LiveHosts         []string
	CrawlDepth        int
	MaxCrawlTime      string
	Threads           int
	ParllelismThreads int
	Headless          bool
	NoSandbox         bool
}

func CrawlHosts(k *KatanaScanner) ([]string, error) {
	katanaArgs := []string{"-jc", "-kf", "-fx", "-xhr", "-jsl", "-aff",
		"-c", strconv.Itoa(k.Threads),
		"-p", strconv.Itoa(k.ParllelismThreads),
		"-d", strconv.Itoa(k.CrawlDepth),
		"-ct", k.MaxCrawlTime,
	}

	if k.Headless {
		katanaArgs = append(katanaArgs, "-headless")
	}
	if k.NoSandbox {
		katanaArgs = append(katanaArgs, "-no-sandbox")
	}
	return utils.RunCommandWithStdinInput("katana", katanaArgs, k.LiveHosts)
}
