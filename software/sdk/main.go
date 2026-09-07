package main

import (
	"os"
	"wr-sdk/theme"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
)

func main() {
	// the event loop
	go func() {
		w := new(app.Window)
		th := theme.Vectorheart() // theme customizes colors, shapes, fonts, etc.
		var ops op.Ops            // records a buffer that tells Gio what to draw and handle and applies them all at once.

		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				os.Exit(0) // clean exit
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)     // the event holds data about the window such as animation timers and size.
				paint.Fill(gtx.Ops, th.Palette.Bg) // color bg

				material.Label(th, 16, "Hello Wristrunners").Layout(gtx) // adds a title.
				e.Frame(gtx.Ops)                                         // draw to screen.
			}
		}
	}()

	app.Main() // keeps the event loop alive
}
