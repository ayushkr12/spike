package katana

import (
	"strconv"

	"github.com/ayushkr12/spike/internal/pkg/utils"
	"github.com/ayushkr12/spike/pkg/config"
)

func CrawlHosts(liveHosts []string, k *config.KatanaConfig) ([]string, error) {
	katanaArgs := []string{"-silent", "-jc", "-kf", "-fx", "-xhr", "-jsl", "-aff",
		"-c", strconv.Itoa(k.Threads),
		"-p", strconv.Itoa(k.ParallelismThreads),
		"-d", strconv.Itoa(k.CrawlDepth),
		"-ct", k.MaxCrawlTime,
	}

	if k.Headless {
		katanaArgs = append(katanaArgs, "-headless")
	}
	if k.NoSandbox {
		katanaArgs = append(katanaArgs, "-no-sandbox")
	}
	return utils.RunCommandWithStdinInput("katana", katanaArgs, liveHosts)
}
