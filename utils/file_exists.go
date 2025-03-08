package utils

import (
	"os"
)

func FileExists(file_path string) bool {
	_, err := os.Stat(file_path)
	if os.IsNotExist(err) {
		return false
	}

	return err == nil
}
