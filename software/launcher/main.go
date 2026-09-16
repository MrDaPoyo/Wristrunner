package main

import (
	"os"
	"time"
	"uikit"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	LoadApps()

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
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return header(gtx, th.Theme) }),
					layout.Flexed(0.4, func(gtx layout.Context) layout.Dimensions { return appGrid(gtx, th.Theme) }),
					layout.Flexed(0.6, func(gtx layout.Context) layout.Dimensions { return alertWidgets(gtx, th.Theme) }),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions { return footer(gtx, th.Theme) }),
				)

				e.Frame(gtx.Ops) // render
			}
		}
	}()

	app.Main() // keeps the event loop alive
}

func header(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
		layout.Flexed(0.33, func(gtx layout.Context) layout.Dimensions {
			return currentTime(gtx, th)
		}),
		layout.Flexed(0.33, func(gtx layout.Context) layout.Dimensions {
			Date := time.Now().Format("Jan 2, 2006")
			return material.Body1(th, Date).Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.Body1(th, "Battery").Layout(gtx)
		}),
	)
}

func appGrid(gtx layout.Context, th *material.Theme) layout.Dimensions {
	app1, app2, app3 := widget.Clickable{}, widget.Clickable{}, widget.Clickable{}
	app4, app5, app6 := widget.Clickable{}, widget.Clickable{}, widget.Clickable{}
	app7, app8, app9 := widget.Clickable{}, widget.Clickable{}, widget.Clickable{}
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return material.Button(th, &app1, "App 1").Layout(gtx)
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return material.Button(th, &app2, "App 2").Layout(gtx)
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return material.Button(th, &app3, "App 3").Layout(gtx)
				}),
			)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return material.Button(th, &app4, "App 4").Layout(gtx)
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return material.Button(th, &app5, "App 5").Layout(gtx)
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return material.Button(th, &app6, "App 6").Layout(gtx)
				}),
			)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return material.Button(th, &app7, "App 7").Layout(gtx)
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return material.Button(th, &app8, "App 8").Layout(gtx)
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return material.Button(th, &app9, "App 9").Layout(gtx)
				}),
			)
		}),
	)
}

func alertWidgets(gtx layout.Context, th *material.Theme) layout.Dimensions {
	//This needs to have a different font and also be centered in it's own little space.
	//Try to find a more optimal way of layouting this.
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return material.Label(th, 16, "W").Layout(gtx)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return material.Label(th, 16, "R").Layout(gtx)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return material.Label(th, 16, "I").Layout(gtx)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return material.Label(th, 16, "S").Layout(gtx)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return material.Label(th, 16, "T").Layout(gtx)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return material.Label(th, 16, "R").Layout(gtx)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return material.Label(th, 16, "U").Layout(gtx)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return material.Label(th, 16, "N").Layout(gtx)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return material.Label(th, 16, "N").Layout(gtx)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return material.Label(th, 16, "E").Layout(gtx)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return material.Label(th, 16, "R").Layout(gtx)
		}),
	)
}

func footer(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return material.Label(th, 14, "Nav / dock").Layout(gtx)
}

func currentTime(gtx layout.Context, th *material.Theme) layout.Dimensions {
	now := time.Now()

	// redraw when the next second begins.
	next := now.Truncate(time.Second).Add(time.Second)
	gtx.Execute(op.InvalidateCmd{At: next})

	return material.Body1(
		th,
		now.Format("15:04:05"),
	).Layout(gtx)
}
