package utils

import (
	"bufio"
	"os"
	"os/exec"
	"regexp"
)

// Check if running on Bedrock
/*
	Example:
	if utils.OnBedrock() {
		// code
	}
*/
func OnBedrock() bool {
	if PathExists("/bedrock") {
		return true
	}

	return false
}

// Count all installed strata's
/*
	Example:
	fmt.Println(utils.Countstrata())
*/
func Countstrata() int {
	if !PathExists("/bedrock") {
		Error("Not running on a Bedrock system.")
		os.Exit(-1)
	}

	strata, _ := os.ReadDir("/bedrock/strata")

	var count = len(strata) - 3

	return count
}

// Detect Bedrock's pmm interface
func Detectpmminterface() *string {
	file, err := os.Open("/bedrock/etc/bedrock.conf")

	if err != nil {
		Error("Failed to read bedrock.conf")
		return nil
	}
	defer file.Close()

	re := regexp.MustCompile(`(?m)^user-interface =(?:"(.*?)"|(.*))$`)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if matches := re.FindStringSubmatch(line); matches != nil {
			if matches[1] != "" {
				return &matches[1]
			}

			if matches[2] != "" {
				return &matches[2]
			}
		}
	}

	return nil
}

// Check what stratum provides the init, useful for checking the currently running stratum.
func Detectstratuminit() string {
	if !PathExists("/bedrock") {
		Error("Not running on a Bedrock system.")
		os.Exit(-1)
	}

	var brl_command = exec.Command("brl", "which", "/sbin/init")

	output, err := brl_command.Output()

	if err != nil {
		Error("Failed to get the init stratum")
	}

	return string(output)
}
