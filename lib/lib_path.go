package lib

import (
	"os"
	"runtime"
	"strings"
)

// UserHomeDir returns the user's home directory
func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	} else if runtime.GOOS == "linux" {
		home := os.Getenv("XDG_CONFIG_HOME")
		if home != "" {
			return home
		}
	}
	return os.Getenv("HOME")
}

// Remove trailing slash if any.
func StripTrailingSlashes(path string) string {
	if strings.HasSuffix(path, "/") {
		if len(path) > 0 && path[len(path)-1] == '/' {
			path = path[0 : len(path)-1]
		}
	} else if strings.HasSuffix(path, "\\") {
		if len(path) > 0 && path[len(path)-1] == '\\' {
			path = path[0 : len(path)-1]
		}
	}
	return path
}
