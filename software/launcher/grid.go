package main

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type AppGrid struct {
	Apps     []App
	Columns  int
	CellMin  unit.Dp
	OnLaunch func(App) // called when a button is clicked

	Clicks []widget.Clickable
}

func NewAppGrid(apps []App, onLaunch func(App)) *AppGrid {
	return &AppGrid{Apps: apps, OnLaunch: onLaunch}
}

func (g *AppGrid) Layout(gtx C, th *material.Theme) D {
	// only 1 Clickable per app
	if len(g.Clicks) < len(g.Apps) {
		g.Clicks = append(g.Clicks, make([]widget.Clickable, len(g.Apps)-len(g.Clicks))...)
	}
	g.Clicks = g.Clicks[:len(g.Apps)]

	// handle clicks
	for i := range g.Clicks {
		for g.Clicks[i].Clicked(gtx) {
			if g.OnLaunch != nil {
				g.OnLaunch(g.Apps[i])
			}
		}
	}

	// if no apps
	if len(g.Apps) == 0 {
		return layout.Center.Layout(gtx, func(gtx C) D {
			return material.Body1(th, "!! No Apps Loaded !!").Layout(gtx)
		})
		// return D{}
	}

	// count columns
	cols := g.Columns
	if cols <= 0 {
		cellMin := g.CellMin
		if cellMin <= 0 {
			cellMin = 150 // width
		}
		cols = gtx.Constraints.Max.X / gtx.Dp(cellMin)
		if cols < 1 {
			cols = 1
		}
	}

	margins := layout.UniformInset(unit.Dp(5))
	if cols > len(g.Apps) {
		cols = len(g.Apps)
	}
	rows := (len(g.Apps) + cols - 1) / cols

	return margins.Layout(gtx, func(gtx C) D {
		rowChildren := make([]layout.FlexChild, rows)
		for r := range rows {
			rowChildren[r] = layout.Flexed(1, func(gtx C) D {
				cellChildren := make([]layout.FlexChild, cols)
				for c := 0; c < cols; c++ {
					i := r*cols + c
					cellChildren[c] = layout.Flexed(1, func(gtx C) D {
						// reserve the space so columns stay aligned
						if i >= len(g.Apps) {
							return D{Size: gtx.Constraints.Max}
						}
						return margins.Layout(gtx, func(gtx C) D {
							return material.Button(th, &g.Clicks[i], g.Apps[i].Name).Layout(gtx)
						})
					})
				}
				return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, cellChildren...)
			})
		}
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx, rowChildren...)
	})
}
