package gau

import (
	"strconv"

	"github.com/ayushkr12/spike/internal/pkg/utils"
)

func GetAllUrls(domain string, Threads int) ([]string, error) {
	gauArgs := []string{domain, "--subs", "--threads", strconv.Itoa(Threads)}
	return utils.RunCommand("gau", gauArgs)
}
