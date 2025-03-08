package utils

import (
	"github.com/hairyhenderson/go-which"
)

func Is_cmd_installed(command string) bool {
	return which.Found(command)
}
