package utils

// Check if running on NixOS
/*
Example:
	if utils.OnNixOS() {
		// code
	}
*/
func OnNixOS() bool {
	if PathExists("/run/current-system/sw/bin") {
		return true
	}

	return false
}
