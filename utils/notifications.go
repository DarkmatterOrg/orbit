package utils

import "github.com/gen2brain/beeep"

// Send an error notification
/*
	Example:
	ErrorNotification("hi", "some text")
*/
func ErrorNotification(title string, text string) {
	beeep.Notify(title, text, "dialog-error")
}

// Send a warning notification
/*
	Example:
	WarnNotification("hi", "some text")
*/
func WarnNotification(title string, text string) {
	beeep.Notify(title, text, "dialog-warning")
}

// Send a info notification
/*
	Example:
	InfoNotification("hi", "some text")
*/
func InfoNotification(title string, text string) {
	beeep.Notify(title, text, "dialog-information")
}

// Send a success notification
/*
	Example:
	DoneNotification("hi", "some text")
*/
func DoneNotification(title string, text string) {
	beeep.Notify(title, text, "dialog-information")
}
