package notes

import (
	"uikit"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func (a *App) layoutList(gtx layout.Context, th *uikit.Theme) layout.Dimensions {
	if len(a.rows) < len(a.Notes) {
		a.rows = append(a.rows, make([]widget.Clickable, len(a.Notes)-len(a.rows))...)
	}
	a.rows = a.rows[:len(a.Notes)]

	for i := range a.rows {
		if a.rows[i].Clicked(gtx) {
			a.openNote(i)
			return a.layoutEdit(gtx, th)
		}
	}
	if a.addBtn.Clicked(gtx) {
		a.newNote()
		return a.layoutEdit(gtx, th)
	}

	inset := layout.UniformInset(unit.Dp(th.Spacing))
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.Button(th.Theme, &a.addBtn, "New note").Layout(gtx)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return a.list.Layout(gtx, len(a.Notes), func(gtx layout.Context, i int) layout.Dimensions {
				return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					label := a.Notes[i].Title
					if label == "" {
						label = "Untitled"
					}
					return material.Button(th.Theme, &a.rows[i], label).Layout(gtx)
				})
			})
		}),
	)
}
