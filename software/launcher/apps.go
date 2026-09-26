package main

import (
	"fmt"
	"lib"
	"os"
	"path"
	"reflect"
	"strings"
)

const MANIFEST_NAME = "manifest.wr" // .wr is such a fire extension
const APP_DIR = "/usr/local/share/wristrunner/apps"

type App struct {
	Name   string // name displayed on the grid.
	Author string // username of who did it.
	Path   string // absolute path to the app's root dir.
}

// every app inside of apps/ is a directory containing a manifest.wr file which is just key: value.
// LoadApps scans apps/ and loads apps with valid manifest.wr files.
func LoadApps() ([]App, error) {
	appDir, err := os.ReadDir(APP_DIR)
	if err != nil {
		return nil, err
	}

	var appList []App

	for _, dir := range appDir {
		if dir.IsDir() {
			manifest, err := os.ReadFile(path.Join(APP_DIR, dir.Name(), MANIFEST_NAME))
			if err == nil { // discard failed reads
				var values = make(map[string]string)

				for line := range strings.Lines(string(manifest)) {
					key, value, found := strings.Cut(line, ": ")
					if found {
						values[lib.Capitalize(key)] = strings.TrimRight(value, "\n") // each line has a newline in the end
					}
				}

				var app App
				assignMatchingFields(&app, values)
				fmt.Printf("Loaded '%s' app.\n", app.Name)
				appList = append(appList, app)
			} else {
				fmt.Println(err)
			}
		}
	}

	return appList, nil
}

// grabs
func assignMatchingFields(app *App, values map[string]string) {
	v := reflect.ValueOf(app).Elem()

	for key, value := range values {
		field := v.FieldByName(key)

		if !field.IsValid() || !field.CanSet() || field.Kind() != reflect.String {
			continue
		}

		field.SetString(value)
	}
}
