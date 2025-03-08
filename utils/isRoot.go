package utils

import "syscall"

func IsRoot() bool {
	if syscall.Getuid() == 0 {
		return true
	}

	return false
}
