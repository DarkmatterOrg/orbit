package utils

import "syscall"

// Check if running as root
/*
Example:
	if utils.IsRoot() {
		// code
	}
*/
func IsRoot() bool {
	if syscall.Getuid() == 0 {
		return true
	}

	return false
}
