package utils

import (
	"time"

	"github.com/briandowns/spinner"
)

var s *spinner.Spinner

/*
Starts a spinner like animation
*/
func StartSpinner() {
	s = spinner.New(spinner.CharSets[43], 150*time.Millisecond)
	s.Start()
}

/*
Stops the animation
*/
func StopSpinner() {
	s.Stop()
}
