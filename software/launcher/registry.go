package main

import (
	"counter"
	"uikit"
	"notes"
)

// when you make a new app, you gotta put it here
var registry = map[string]func() uikit.App{
	"Counter": counter.New,
	"Notes":   notes.New,
}
