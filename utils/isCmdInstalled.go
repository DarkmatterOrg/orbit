package utils

import (
	"github.com/hairyhenderson/go-which"
)

// Check if a comamnd is installed
/*
	Example:
	if IsCmdInstalled("firefox") {
		// code
	}
*/
func IsCmdInstalled(command string) bool {
	return which.Found(command)
}
