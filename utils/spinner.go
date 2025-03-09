package utils

import (
	"time"

	"github.com/briandowns/spinner"
)

var s *spinner.Spinner

func StartSpinner() {
	s = spinner.New(spinner.CharSets[43], 150*time.Millisecond)
	s.Start()
}

func StopSpinner() {
	s.Stop()
}
