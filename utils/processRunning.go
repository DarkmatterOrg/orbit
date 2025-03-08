package utils

import (
	"os/exec"
	"strings"
)

func IsProcessRunning(process string) bool {
	cmd := exec.Command("pgrep", "-x", process)
	output, err := cmd.Output()
	return err == nil && strings.TrimSpace(string(output)) != ""
}
