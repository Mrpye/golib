// Package path provides functions for manipulating file paths
package path

import (
	"os"
	"runtime"
	"strings"
)

// If the OS is Windows, return the HOMEDRIVE and HOMEPATH environment variables concatenated together.
// If the OS is Linux, return the XDG_CONFIG_HOME environment variable. Otherwise, return the HOME
// environment variable
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

// StripTrailingSlashes If the path ends with a slash, remove it
// - path: the path to strip
// - returns: the path with the trailing slash removed
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
