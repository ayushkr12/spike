package httpx

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/ayushkr12/spike/internal/pkg/utils"
)

func Probe(Hosts []string) ([]string, error) {
	var stdout strings.Builder // Create a new strings.Builder to capture the output
	cmd := exec.Command("httpx", "-s")
	cmd.Stdout = &stdout
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdin pipe: %w", err)
	}
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start httpx command: %w", err)
	}
	_, err = stdin.Write([]byte(utils.JoinSlice(Hosts)))
	if err != nil {
		return nil, fmt.Errorf("failed to write to stdin: %w", err)
	}
	if err := stdin.Close(); err != nil {
		return nil, fmt.Errorf("failed to close stdin: %w", err)
	}
	if err := cmd.Wait(); err != nil {
		return nil, fmt.Errorf("httpx command failed: %w", err)
	}
	lines := utils.LinesToSlice(stdout.String())              // Split the output into lines and trim whitespace
	liveHosts := utils.RemoveDuplicatesAndEmptyStrings(lines) // Remove duplicates and empty strings
	return liveHosts, nil
}
