package main

import (
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

	// the event loop
	go func() {
		w := new(app.Window)
		th := uikit.VectorheartTheme() // theme customizes colors, shapes, fonts, etc.
		var ops op.Ops                 // records a buffer that tells Gio what to draw and handle and applies them all at once.

		grid := NewAppGrid(apps, func(a App) {
			log.Println("launching", a.Name, "from", a.Path)
		})

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
					layout.Rigid(func(gtx C) D { return header(gtx, th.Theme) }),
					layout.Flexed(0.4, func(gtx C) D {
						return grid.Layout(gtx, th.Theme)
					}),
					layout.Flexed(0.6, func(gtx C) D { return alertWidgets(gtx, th.Theme) }),
					layout.Rigid(func(gtx C) D { return footer(gtx, th.Theme) }),
				)

				e.Frame(gtx.Ops) // render
			}
		}
	}()

	app.Main() // keeps the event loop alive
}

func header(gtx C, th *material.Theme) D {
	margins := layout.Inset{
		Top:    unit.Dp(5),
		Bottom: unit.Dp(0),
		Left:   unit.Dp(5),
		Right:  unit.Dp(5),
	}
	batteryMargins := layout.Inset{
		Top:    unit.Dp(5),
		Bottom: unit.Dp(0),
		Left:   unit.Dp(55),
		Right:  unit.Dp(5),
	}
	return margins.Layout(gtx, func(gtx C) D {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				return margins.Layout(gtx, func(gtx C) D {
					return currentTime(gtx, th)
				})
			}),
			layout.Rigid(func(gtx C) D {
				return margins.Layout(gtx, func(gtx C) D {
					return material.Label(th, 16, "/").Layout(gtx)
				})
			}),
			layout.Rigid(func(gtx C) D {
				Date := time.Now().Format("Jan 2, 2006")
				return margins.Layout(gtx, func(gtx C) D {
					return material.Body1(th, Date).Layout(gtx)
				})
			}),
			layout.Rigid(func(gtx C) D {
				return batteryMargins.Layout(gtx, func(gtx C) D {
					return material.Body1(th, "Battery").Layout(gtx)
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
	return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return margins.Layout(gtx, func(gtx C) D {
				return material.Label(th, 16, "W").Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx C) D {
			return margins.Layout(gtx, func(gtx C) D {
				return material.Label(th, 16, " R").Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx C) D {
			return margins.Layout(gtx, func(gtx C) D {
				return material.Label(th, 16, "  I").Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx C) D {
			return margins.Layout(gtx, func(gtx C) D {
				return material.Label(th, 16, " S").Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx C) D {
			return margins.Layout(gtx, func(gtx C) D {
				return material.Label(th, 16, " T").Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx C) D {
			return margins.Layout(gtx, func(gtx C) D {
				return material.Label(th, 16, " R").Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx C) D {
			return margins.Layout(gtx, func(gtx C) D {
				return material.Label(th, 16, " U").Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx C) D {
			return margins.Layout(gtx, func(gtx C) D {
				return material.Label(th, 16, " N").Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx C) D {
			return margins.Layout(gtx, func(gtx C) D {
				return material.Label(th, 16, " N").Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx C) D {
			return margins.Layout(gtx, func(gtx C) D {
				return material.Label(th, 16, " E").Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx C) D {
			return margins.Layout(gtx, func(gtx C) D {
				return material.Label(th, 16, " R").Layout(gtx)
			})
		}),
	)
}

func footer(gtx C, th *material.Theme) D {
	return material.Label(th, 14, "Nav / dock").Layout(gtx)
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
