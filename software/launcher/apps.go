package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

const MANIFEST_NAME = "manifest.wr" // .wr is such a fire extension
const APP_DIR = "../apps"

type App struct {
	Name   string // name displayed on the grid.
	Author string // username of who did it.
	Path   string // absolute path to the app's root dir.
}

// every app inside of apps/ is a directory containing a manifest.wr file which is just key: value.
// LoadApps scans apps/ and loads apps with valid manifest.wr files.
func LoadApps() error {
	appDir, err := os.ReadDir(APP_DIR)
	if err != nil {
		return err
	}

	for _, dir := range appDir {
		if dir.IsDir() {
			manifest, err := os.ReadFile(path.Join(APP_DIR, dir.Name(), MANIFEST_NAME))
			if err == nil { // discard failed reads
				// todo: parse manifest
				for line := range strings.Lines(string(manifest)) {
					vars := strings.Split(line, ": ")
					key := vars[0]
					value := strings.TrimRight(vars[1], "\n") // each line has a newline in the end
					fmt.Println(key + value)
				}
			} else {
				fmt.Println(err)
			}

		}
	}

	return nil
}
