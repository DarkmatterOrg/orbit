package utils

import "os"

// Append a string to a file
/*
	Example
	AppendFile("Hello, world!", "test.txt")
*/
func AppendFile(content string, filepath string) error {
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.WriteString(content); err != nil {
		return err
	}

	return nil
}
