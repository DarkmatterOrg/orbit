package utils

import (
	"bufio"
	"os"
	"regexp"
)

// Detect the OS by checking the os-release file
func Getos() string {
	os_release_path := "/etc/os-release"

	if OnBedrock() {
		os_release_path = "/bedrock/etc/os-release"
	}

	var system_name = readOSRelease(os_release_path)

	return *system_name
}

// Using regex to read only the part after PRETTY_NAME
func readOSRelease(path string) *string {
	file, err := os.Open(path)

	if err != nil {
		Error("Failed to read os-release")
		return nil
	}
	defer file.Close()

	re := regexp.MustCompile(`(?m)^PRETTY_NAME=(?:"(.*?)"|(.*))$`)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if matches := re.FindStringSubmatch(line); matches != nil {
			if matches[1] != "" {
				return &matches[1]
			} else if matches[2] != "" {
				return &matches[2]
			}
		}
	}

	return nil
}
