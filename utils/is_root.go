package utils

import "syscall"

func Is_root() bool {
	if syscall.Getuid() == 0 {
		return true
	}

	return false
}
