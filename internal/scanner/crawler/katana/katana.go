package katana

import (
	"strconv"

	"github.com/ayushkr12/spike/internal/pkg/utils"
)

func CrawlHosts(liveHosts []string, Threads int, ParllelismThreads int, Headless bool) ([]string, error) {
	katanaArgs := []string{"-s", "-jc", "-kf", "-fx", "-xhr", "-jsl", "-aff",
		"-t", strconv.Itoa(Threads), "-p", strconv.Itoa(ParllelismThreads)}
	if Headless {
		katanaArgs = append(katanaArgs, "-headless")
	}
	return utils.RunCommandWithStdinInput("katana", katanaArgs, liveHosts)
}
