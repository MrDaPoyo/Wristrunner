package uikit

import "gioui.org/layout"

type AppInfo struct {
	Title  string
	Author string
}

type App interface {
	Layout(gtx layout.Context, th *Theme) layout.Dimensions // called when the app is active
	Close()                                                 // called when the app is closedç
}
