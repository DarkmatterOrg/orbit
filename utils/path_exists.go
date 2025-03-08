package utils

import (
	"os"
)

func Pathexists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
