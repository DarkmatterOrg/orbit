package utils

import (
	"os"
	"os/exec"
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
