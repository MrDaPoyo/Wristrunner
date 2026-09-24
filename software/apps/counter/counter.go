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
	dec   widget.Clickable
}

func New() uikit.App { return &App{} }

func (a *App) Layout(gtx layout.Context, th *uikit.Theme) layout.Dimensions {
	if a.inc.Clicked(gtx) {
		a.count++
	}
	if a.dec.Clicked(gtx) {
		a.count--
	}

	return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle, Spacing: layout.SpaceBetween, Gap: th.Spacing}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return material.Body1(th.Theme, strconv.Itoa(a.count)).Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return material.Button(th.Theme, &a.inc, "+1").Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return material.Button(th.Theme, &a.dec, "-1").Layout(gtx)
			}))
	})
}

func (a *App) Close() {}
