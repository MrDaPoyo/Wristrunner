package main

import (
	"uikit"

	"gioui.org/widget"
)

type Shell struct {
	grid    *AppGrid
	active  string // string acts as the id
	running map[string]uikit.App

	homeBtn  widget.Clickable
	closeBtn widget.Clickable
}

func NewShell(apps []App) *Shell {
	s := &Shell{running: map[string]uikit.App{}}
	s.grid = NewAppGrid(apps, s.Open)
	return s
}

// opens/focuses an app
func (s *Shell) Open(a App) {
	if _, ok := s.running[a.Name]; !ok {
		ctor, ok := registry[a.Name]
		if !ok {
			return // there's no constructor/entrypoint
		}
		s.running[a.Name] = ctor() // ctor is the constructor
	}

	s.active = a.Name // focus
}

func (s *Shell) Home() { s.active = "" } // empty means no app

func (s *Shell) Active() string { return s.active }

func (s *Shell) Layout(gtx C, th *uikit.Theme) D {
	if s.active != "" {
		if a, ok := s.running[s.active]; ok {
			return a.Layout(gtx, th)
		}
	}
	return s.grid.Layout(gtx, th.Theme)
}

// "closes" the app, but in reality it just hides it away
func (s *Shell) Close() {
	if s.active == "" {
		return
	}
	if a, ok := s.running[s.active]; ok {
		a.Close()
		delete(s.running, s.active)
	}
	s.active = ""
}
