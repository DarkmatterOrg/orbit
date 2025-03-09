package utils

import (
	"fmt"

	"github.com/fatih/color"
)

var notice = color.New(color.FgCyan, color.Bold).SprintFunc()
var info = color.New(color.FgBlue, color.Bold).SprintFunc()
var error = color.New(color.FgRed, color.Bold).SprintFunc()
var warning = color.New(color.FgYellow, color.Bold).SprintFunc()
var done = color.New(color.FgGreen, color.Bold).SprintFunc()
var debug = color.New(color.FgHiCyan, color.Bold).SprintFunc()

func Notice(msg string) {
	fmt.Println(notice("NOTICE") + ": " + msg)
}

func Noticef(msg string) string {
	return fmt.Sprintf(notice("NOTICE") + ": " + msg)
}

func Info(msg string) {

	fmt.Println(info("INFO") + ": " + msg)
}

func Infof(msg string) string {
	return fmt.Sprintf(info("INFO") + ": " + msg)
}

func Error(msg string) {
	fmt.Println(error("ERR") + ": " + msg)
}

func Errorf(msg string) string {
	return fmt.Sprintf(error("ERROR") + ": " + msg)
}

func Warn(msg string) {
	fmt.Println(warning("WARN") + ": " + msg)
}

func Warnf(msg string) string {
	return fmt.Sprintf(warning("WARN") + ": " + msg)
}

func Done(msg string) {
	fmt.Println(done("DONE") + ": " + msg)
}

func Donef(msg string) string {
	return fmt.Sprintf(done("DONE") + ": " + msg)
}

func Debug(msg string) {
	fmt.Println(debug("DEBUG") + ": " + msg)
}

func Debugf(msg string) string {
	return fmt.Sprintf(debug("DEBUG") + ": " + msg)
}
