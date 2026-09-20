package counter

import (
	"strconv"
	"uikit"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type App struct {
	count int
	inc   widget.Clickable
}

func New() uikit.App { return &App{} }

func (a *App) Layout(gtx layout.Context, th *uikit.Theme) layout.Dimensions {
	if a.inc.Clicked(gtx) {
		a.count++
	}

	return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return material.Body1(th.Theme, "count = "+strconv.Itoa(a.count)).Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return material.Button(th.Theme, &a.inc, "increment").Layout(gtx)
			}))
	})
}

func (a *App) Close() {}
