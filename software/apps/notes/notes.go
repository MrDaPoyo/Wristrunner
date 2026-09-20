package notes

import (
	"time"
	"uikit"

	"gioui.org/layout"
	"gioui.org/widget/material"
)

type Note struct {
	Title      string
	Content    string
	CreatedAt  time.Time
	ModifiedAt time.Time
}

type App struct {
	Notes []Note
}

func New() uikit.App { return &App{} }

func (a *App) Layout(gtx layout.Context, th *uikit.Theme) layout.Dimensions {
	return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return material.Body1(th.Theme, "Notes here").Layout(gtx)
	})
}

func (a *App) Close() {}
