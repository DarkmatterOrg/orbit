package utils

import (
	"fmt"
	"os/exec"
)

func Runcommand(command string, arg string) {
	fullCommand := fmt.Sprintf("%s %s", command, arg)

	cmd := exec.Command("/bin/sh", "-c", fullCommand)
	err := cmd.Run()

	if err != nil {
		fmt.Println("Failed to run the command")
	}
}
