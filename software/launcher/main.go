package main

import (
	"os"
	"time"
	"wr-ui/theme"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
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

				layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return header(gtx, th) }),
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions { return body(gtx, th) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return footer(gtx, th) }),
				)

				e.Frame(gtx.Ops) // render                                       // draw to screen.
			}
		}
	}()

	app.Main() // keeps the event loop alive
}

func header(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return layout.Inset{
		Top:    unit.Dp(8),
		Bottom: unit.Dp(8),
		Left:   unit.Dp(16),
		Right:  unit.Dp(16),
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {

		return layout.Flex{
			Axis: layout.Horizontal,
		}.Layout(gtx,

			// left third
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return layout.W.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return material.Body1(th, "Notifications").Layout(gtx)
				})
			}),

			// center third
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return currentTime(gtx, th)
				})
			}),

			// right third
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return layout.E.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return material.Body1(th, "Stats").Layout(gtx)
				})
			}),
		)
	})
}

func body(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return material.Label(th, 16, "App grid goes here").Layout(gtx)
}

func footer(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return material.Label(th, 14, "Nav / dock").Layout(gtx)
}

func currentTime(gtx layout.Context, th *material.Theme) layout.Dimensions {
	now := time.Now()

	// redraw when the next second begins.
	next := now.Truncate(time.Second).Add(time.Second)

	gtx.Execute(op.InvalidateCmd{
		At: next,
	})

	return material.Body1(
		th,
		now.Format("15:04:05"),
	).Layout(gtx)
}
