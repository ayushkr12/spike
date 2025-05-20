package httpx

import (
	"strconv"

	"github.com/ayushkr12/spike/internal/pkg/utils"
)

func Probe(Hosts []string, Threads int) ([]string, error) {
	return utils.RunCommandWithStdinInput("httpx", []string{"-s", "-t", strconv.Itoa(Threads)}, Hosts)
}
