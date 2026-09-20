package lib

import (
	"os"
	"path/filepath"
	"strings"
)

// returns the path to the directory for an app to store data in
func DataDir(appID string) (string, error) {
	base := os.Getenv("XDG_DATA_HOME") // linux-specific, but it's fine and it'll work on windows too
	if base == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		base = filepath.Join(home, ".local", "share")
	}

	dir := filepath.Join(base, "wristrunner", strings.ToLower(appID))
	if err := os.MkdirAll(dir, 0o700); err != nil { // 0700 = read, write, execute
		return "", err
	}

	return dir, nil
}
