package utils

import (
	"fmt"

	"github.com/fatih/color"
)

var notice_msg = color.New(color.FgCyan, color.Bold).SprintFunc()
var info_msg = color.New(color.FgBlue, color.Bold).SprintFunc()
var error_msg = color.New(color.FgRed, color.Bold).SprintFunc()
var warning_msg = color.New(color.FgYellow, color.Bold).SprintFunc()
var done_msg = color.New(color.FgGreen, color.Bold).SprintFunc()
var debug_msg = color.New(color.FgHiCyan, color.Bold).SprintFunc()

func Notice(msg string) {
	fmt.Println(notice_msg("NOTICE") + ": " + msg)
}

func Noticef(msg string) string {
	return fmt.Sprintf(notice_msg("NOTICE") + ": " + msg)
}

func Info(msg string) {
	fmt.Println(info_msg("INFO") + ": " + msg)
}

func Infof(msg string) string {
	return fmt.Sprintf(info_msg("INFO") + ": " + msg)
}

func Error(msg string) {
	fmt.Println(error_msg("ERR") + ": " + msg)
}

func Errorf(msg string) string {
	return fmt.Sprintf(error_msg("ERROR") + ": " + msg)
}

func Warn(msg string) {
	fmt.Println(warning_msg("WARN") + ": " + msg)
}

func Warnf(msg string) string {
	return fmt.Sprintf(warning_msg("WARN") + ": " + msg)
}

func Done(msg string) {
	fmt.Println(done_msg("DONE") + ": " + msg)
}

func Donef(msg string) string {
	return fmt.Sprintf(done_msg("DONE") + ": " + msg)
}

func Debug(msg string) {
	fmt.Println(debug_msg("DEBUG") + ": " + msg)
}

func Debugf(msg string) string {
	return fmt.Sprintf(debug_msg("DEBUG") + ": " + msg)
}
