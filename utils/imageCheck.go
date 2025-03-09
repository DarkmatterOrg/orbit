package utils

import (
	"errors"
	"os"
	"strings"
)

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

func IsCurrentImage(imageName string) (bool, error) {
	imageType := Getimagetype()
	if imageType == "" {
		return false, errors.New("Couldn't get image type")
	}

	return strings.Contains(imageType, strings.ToLower(imageName)), nil
}
