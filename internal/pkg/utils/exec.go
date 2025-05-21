// internal/pkg/utils/exec.go
package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	log "github.com/ayushkr12/logger"
)

// RunCommandWithStdinInput runs a command with StdinInput piped to stdin,
// and returns stdout lines with duplicates and empty lines removed.
func RunCommandWithStdinInput(cmdName string, args []string, StdinInput []string) ([]string, error) {
	var stdout bytes.Buffer
	cmd := exec.Command(cmdName, args...)
	cmd.Stdout = &stdout

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to get stdin: %w", err)
	}

	log.Debugf("Running command: %s", cmd.String())

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start command: %w", err)
	}

	_, err = stdin.Write([]byte(strings.Join(StdinInput, "\n")))
	if err != nil {
		return nil, fmt.Errorf("failed to write to stdin: %w", err)
	}

	if err := stdin.Close(); err != nil {
		return nil, fmt.Errorf("failed to close stdin: %w", err)
	}

	if err := cmd.Wait(); err != nil {
		return nil, fmt.Errorf("command failed: %w", err)
	}

	lines := LinesToSlice(stdout.String())
	return RemoveDuplicatesAndEmptyStrings(lines), nil
}
