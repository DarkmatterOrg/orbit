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

// Write a notice message
/*
	Example:
	Notice("You forgot to turn off the dishwasher!")
*/
func Notice(msg string) {
	fmt.Println(noticeMsg("NOTICE") + ": " + msg)
}

// Write a notice message
/*
	Example:
	Noticef("You forgot to turn off the dishwasher!")
*/
func Noticef(msg string) string {
	return fmt.Sprintf(noticeMsg("NOTICE") + ": " + msg)
}

// Write a info message
/*
	Example:
	Info("Today is Sunday!")
*/
func Info(msg string) {
	fmt.Println(infoMsg("INFO") + ": " + msg)
}

// Write a info message
/*
	Example:
	Infof("The weather is cloudy")
*/
func Infof(msg string) string {
	return fmt.Sprintf(infoMsg("INFO") + ": " + msg)
}

// Write an error message
/*
	Example:
	Error("Access denied")
*/
func Error(msg string) {
	fmt.Println(errorMsg("ERR") + ": " + msg)
}

// Write an error message
/*
	Example:
	Errorf("No such file or directory")
*/
func Errorf(msg string) string {
	return fmt.Sprintf(errorMsg("ERROR") + ": " + msg)
}

// Write a warn message
/*
	Example:
	Warn("Timezone is incorrect!")
*/
func Warn(msg string) {
	fmt.Println(warningMsg("WARN") + ": " + msg)
}

// Write a warn message
/*
	Example:
	Warnf("Timezone is incorrect!")
*/
func Warnf(msg string) string {
	return fmt.Sprintf(warningMsg("WARN") + ": " + msg)
}

// Write a success message
/*
	Example:
	Done("Download finished")
*/
func Done(msg string) {
	fmt.Println(doneMsg("DONE") + ": " + msg)
}

// Write a success message
/*
	Example:
	Donef("Download finished")
*/
func Donef(msg string) string {
	return fmt.Sprintf(doneMsg("DONE") + ": " + msg)
}

// Write a debug message
/*
	Example:
	Debug("Process 1 is still running")
*/
func Debug(msg string) {
	fmt.Println(debugMsg("DEBUG") + ": " + msg)
}

// Write a debug message
/*
	Example:
	Debugf("Process 1 is still running")
*/
func Debugf(msg string) string {
	return fmt.Sprintf(debugMsg("DEBUG") + ": " + msg)
}
