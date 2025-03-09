package utils

import (
	"fmt"

	"github.com/fatih/color"
)

var noticeMsg = color.New(color.FgCyan, color.Bold).SprintFunc()
var infoMsg = color.New(color.FgBlue, color.Bold).SprintFunc()
var errorMsg = color.New(color.FgRed, color.Bold).SprintFunc()
var warningMsg = color.New(color.FgYellow, color.Bold).SprintFunc()
var doneMsg = color.New(color.FgGreen, color.Bold).SprintFunc()
var debugMsg = color.New(color.FgHiCyan, color.Bold).SprintFunc()

func Notice(msg string) {
	fmt.Println(noticeMsg("NOTICE") + ": " + msg)
}

func Noticef(msg string) string {
	return fmt.Sprintf(noticeMsg("NOTICE") + ": " + msg)
}

func Info(msg string) {
	fmt.Println(infoMsg("INFO") + ": " + msg)
}

func Infof(msg string) string {
	return fmt.Sprintf(infoMsg("INFO") + ": " + msg)
}

func Error(msg string) {
	fmt.Println(errorMsg("ERR") + ": " + msg)
}

func Errorf(msg string) string {
	return fmt.Sprintf(errorMsg("ERROR") + ": " + msg)
}

func Warn(msg string) {
	fmt.Println(warningMsg("WARN") + ": " + msg)
}

func Warnf(msg string) string {
	return fmt.Sprintf(warningMsg("WARN") + ": " + msg)
}

func Done(msg string) {
	fmt.Println(doneMsg("DONE") + ": " + msg)
}

func Donef(msg string) string {
	return fmt.Sprintf(doneMsg("DONE") + ": " + msg)
}

func Debug(msg string) {
	fmt.Println(debugMsg("DEBUG") + ": " + msg)
}

func Debugf(msg string) string {
	return fmt.Sprintf(debugMsg("DEBUG") + ": " + msg)
}
