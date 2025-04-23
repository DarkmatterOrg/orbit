package utils

import (
	"os"
	"strings"
)

// Retunrs the type of image for Darkmatterorg images
func Getimagetype() string {
	paths := []string{
		"/usr/share/horizon/image_type",
		"/usr/share/nova/image_type",
		"/usr/share/umbra/image_type",
	}

	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	return ""
}

// Checks if the imagename is the current running image
/*
	Example:
	if IsCurrentImage("aster") {
		// code
	}
*/
func IsCurrentImage(imageName string) bool {
	imageType := Getimagetype()
	if imageType == "" {
		return false
	}

	return strings.Contains(imageType, strings.ToLower(imageName))
}
