package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func Notice(msg string) {
	notice := color.New(color.FgCyan, color.Bold).SprintFunc()

	fmt.Println(notice("NOTICE") + ": " + msg)
}

func Error(msg string) {
	error := color.New(color.FgRed, color.Bold).SprintFunc()

	fmt.Println(error("ERR") + ": " + msg)
}

func Warn(msg string) {
	warning := color.New(color.FgYellow, color.Bold).SprintFunc()

	fmt.Println(warning("WARN") + ": " + msg)
}

func Done(msg string) {
	done := color.New(color.FgGreen, color.Bold).SprintFunc()

	fmt.Println(done("DONE") + ": " + msg)
}
