package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"uikit"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"

	"github.com/distatus/battery"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

func main() {
	apps, err := LoadApps()
	if err != nil {
		log.Panic(err)
	}

	shell := NewShell(apps)
	battery, err := battery.Get(0)
	if err != nil {
		fmt.Println("Could not get battery info!")
	}

	// the event loop
	go func() {
		w := new(app.Window)
		th := uikit.VectorheartTheme() // theme customizes colors, shapes, fonts, etc.
		var ops op.Ops                 // records a buffer that tells Gio what to draw and handle and applies them all at once.

		w.Option(
			app.Title("Wristrunner Launcher"),
			app.Size(unit.Dp(450), unit.Dp(800)),    // 16:9
			app.MinSize(unit.Dp(450), unit.Dp(800)), // 16:9
			app.MaxSize(unit.Dp(450), unit.Dp(800)), // 16:9
		)
		for {
			switch e := w.Event().(type) {
			case app.DestroyEvent:
				os.Exit(0) // clean exit
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)     // the event holds data about the window such as animation timers and size.
				paint.Fill(gtx.Ops, th.Palette.Bg) // color bg

				layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx C) D { return header(gtx, th.Theme, battery) }),
					layout.Flexed(0.4, func(gtx C) D {
						return shell.Layout(gtx, th)
					}),
					layout.Flexed(0.6, func(gtx C) D { return alertWidgets(gtx, th.Theme) }),
					layout.Rigid(func(gtx C) D { return footer(gtx, th.Theme, shell) }),
				)

				e.Frame(gtx.Ops) // render
			}
		}
	}()

	app.Main() // keeps the event loop alive
}

func header(gtx C, th *material.Theme, battery *battery.Battery) D {
	margins := layout.Inset{
		Top:    unit.Dp(5),
		Bottom: unit.Dp(0),
		Left:   unit.Dp(5),
		Right:  unit.Dp(5),
	}
	batteryMargins := layout.Inset{
		Top:    unit.Dp(5),
		Bottom: unit.Dp(0),
		Left:   unit.Dp(5),
		Right:  unit.Dp(5),
	}

	return margins.Layout(gtx, func(gtx C) D {
		return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				return margins.Layout(gtx, func(gtx C) D {
					return currentTime(gtx, th)
				})
			}),
			// layout.Rigid(func(gtx C) D {
			// 	return margins.Layout(gtx, func(gtx C) D {
			// 		return material.Label(th, 16, "/").Layout(gtx)
			// 	})
			// }),
			layout.Rigid(func(gtx C) D {
				Date := time.Now().Format("Jan 2, 2006")
				return margins.Layout(gtx, func(gtx C) D {
					return material.Body1(th, Date).Layout(gtx)
				})
			}),
			layout.Rigid(func(gtx C) D {
				return batteryMargins.Layout(gtx, func(gtx C) D {
					percentage := battery.Current / battery.Full * 100
					return material.Body1(th, fmt.Sprintf("%.0f%%", percentage)).Layout(gtx)
				})
			}),
		)
	})
}

func alertWidgets(gtx C, th *material.Theme) D {
	margins := layout.Inset{
		Top:    unit.Dp(5),
		Left:   unit.Dp(10),
		Right:  unit.Dp(10),
		Bottom: unit.Dp(12),
	}

	children := make([]layout.FlexChild, 0, len("WRISTRUNNER"))

	for _, char := range "WRISTRUNNER" {
		children = append(children, layout.Rigid(func(gtx C) D {
			return margins.Layout(gtx, func(gtx C) D {
				return material.Label(th, 16, string(char)).Layout(gtx)
			})
		}))
	}
	return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx)
}

func footer(gtx C, th *material.Theme, shell *Shell) D {
	running := shell.Active() != ""

	if running && shell.homeBtn.Clicked(gtx) {
		shell.Home()
	}
	if running && shell.closeBtn.Clicked(gtx) {
		shell.Close()
	}

	home := material.Button(th, &shell.homeBtn, "Home")
	close := material.Button(th, &shell.closeBtn, "Close")
	if !running { // greyed out on the grid
		home.Background = th.Palette.ContrastFg
		home.Color = th.Palette.Fg
		close.Background = th.Palette.ContrastFg
		close.Color = th.Palette.Fg
	}

	return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceEvenly}.Layout(gtx,
		layout.Rigid(func(gtx C) D { return home.Layout(gtx) }),
		layout.Rigid(func(gtx C) D { return close.Layout(gtx) }),
	)
}

func currentTime(gtx C, th *material.Theme) D {
	now := time.Now()

	// redraw when the next second begins.
	next := now.Truncate(time.Second).Add(time.Second)
	gtx.Execute(op.InvalidateCmd{At: next})

	return material.Body1(
		th,
		now.Format("15:04:05"),
	).Layout(gtx)
}
