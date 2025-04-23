package utils

import (
	"os"
)

// Write a string to file
/*
	Example:
	WriteFile("test.txt", "Hello, Go!")
*/
func WriteFile(fullFilePath string, stringToWrite string) error {
	file, err := os.Create(fullFilePath)
	if err != nil {
		return err
	}

	_, err = file.WriteString(stringToWrite)
	if err != nil {
		file.Close()
		return err
	}

	file.Close()
	return nil
}
