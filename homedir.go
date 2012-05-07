// Package homedir finds the user's home directory.
package homedir

import (
	"os/user"
	"path"
	"runtime"
)

// Default Home directory for the current user
func Get() string {
	user, err := user.Current()
	if err != nil {
		return "/"
	}
	username := user.Username

	switch runtime.GOOS {
	case "darwin":
		return path.Join("/Users", username)
	}
	return path.Join("/home", username)
}
