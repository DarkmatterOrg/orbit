package utils

import (
	"github.com/hairyhenderson/go-which"
)

func IsCmdInstalled(command string) bool {
	return which.Found(command)
}
