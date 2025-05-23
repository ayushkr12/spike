package subfinder

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/ayushkr12/spike/internal/pkg/utils"
)

// GetSubdomains runs the subfinder command to find subdomains for a given domain.
func GetSubdomains(domain string, Threads int) ([]string, error) {
	var stdout strings.Builder // Create a new strings.Builder to capture the output
	cmd := exec.Command("subfinder", "-d", domain, "-silent", "-t", strconv.Itoa(Threads), "-all")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	// split the output into lines and trim whitespace
	lines := utils.LinesToSlice(stdout.String())
	lines = append(lines, domain)                              // Append the root domain to the list of subdomains
	subdomains := utils.RemoveDuplicatesAndEmptyStrings(lines) // Remove duplicates and empty strings
	return subdomains, nil
}
