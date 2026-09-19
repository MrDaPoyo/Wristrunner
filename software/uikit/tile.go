package uikit

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// tappable icon used by the launcher's grid.
type AppTile struct {
	Click widget.Clickable
	Icon  widget.Image // must be square
	Label string       // fallback/depends on the launcher's grid
}

func (t *AppTile) Layout(gtx layout.Context, th *Theme) layout.Dimensions {
	return t.Click.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Dp(th.Spacing)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					sq := t.Icon.Layout(gtx)
					return sq
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					lbl := material.Body2(th.Theme, t.Label)
					lbl.Color = th.Palette.Fg
					return lbl.Layout(gtx)
				}),
			)
		})
	})
}
