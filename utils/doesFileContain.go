package utils

import (
	"os"
	"regexp"
)

func DoesFileContain(str string, filepath string) bool {
	b, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	isExist, err := regexp.Match(str, b)
	if err != nil {
		panic(err)
	}
	return isExist
}