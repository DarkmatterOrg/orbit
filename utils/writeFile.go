package utils

import (
	"os"
)

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
