package utils

import (
	"os"
	"strings"
)

func DoesFileContain(str string, filepath string) bool {
	b, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	s := string(b)
	return strings.Contains(s, str)
}