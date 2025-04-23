package utils

import (
	"os"
	"path/filepath"
	"strings"
)

// Checks if a process is running
/*
	Example:
	IsProcessRunning("firefox")
*/
func IsProcessRunning(target string) bool {
	procPath := "/proc"

	entries, err := os.ReadDir(procPath)
	if err != nil {
		return false
	}

	for _, entry := range entries {
		if !entry.IsDir() || !isNumeric(entry.Name()) {
			continue
		}

		commPath := filepath.Join(procPath, entry.Name(), "comm")
		data, err := os.ReadFile(commPath)
		if err != nil {
			continue
		}

		name := strings.TrimSpace(string(data))
		if name == target {
			return true
		}
	}

	return false
}

func isNumeric(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
