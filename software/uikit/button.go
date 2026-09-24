package uikit

import (
	"fmt"
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Button struct {
	Clickable widget.Clickable
	Action    func()

	Text       string
	Background color.NRGBA
	Foreground color.NRGBA
}

const RADIUS = 5

func (b *Button) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return b.Clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Stack{}.Layout(gtx,
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				size := gtx.Constraints.Min

				r := clip.RRect{
					Rect: image.Rectangle{
						Max: size,
					},
					NE: gtx.Dp(RADIUS),
					NW: gtx.Dp(RADIUS),
					SE: gtx.Dp(RADIUS),
					SW: gtx.Dp(RADIUS),
				}

				if b.Clickable.Clicked(gtx) {
					b.Action()
					fmt.Println("lol")
				}

				defer r.Push(gtx.Ops).Pop()

				paint.Fill(gtx.Ops, b.Background)

				return layout.Dimensions{
					Size: size,
				}
			}),

			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(12)).Layout(
					gtx,
					func(gtx layout.Context) layout.Dimensions {
						label := material.Body1(th, b.Text)
						label.Color = b.Foreground
						return label.Layout(gtx)
					},
				)
			}),
		)
	})
}
