package utils

import (
	"os"
	"strings"
)

// Checks if a file contains a string
/*
	Example:
	if DoesFileContain("hello", "test.txt") {
		// code
	}
*/
func DoesFileContain(str string, filepath string) bool {
	b, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	s := string(b)
	return strings.Contains(s, str)
}
