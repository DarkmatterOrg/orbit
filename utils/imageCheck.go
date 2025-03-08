package utils

import "os"

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
