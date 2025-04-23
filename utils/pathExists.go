package utils

import (
	"os"
)

// Check if path exists
/*
	Example:
	if PathExists("/etc/os-release") {
		// code
	}
*/
func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}
