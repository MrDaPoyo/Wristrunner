package notes

import (
	"log"
	"time"
	"uikit"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func (a *App) layoutEdit(gtx layout.Context, th *uikit.Theme) layout.Dimensions {
	save := a.saveBtn.Clicked(gtx)
	if a.backBtn.Clicked(gtx) {
		a.screen = screenList // discard: editors are only a working copy
		return a.layoutList(gtx, th)
	}

	inset := layout.UniformInset(unit.Dp(th.Spacing))

	dims := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.Button(th.Theme, &a.backBtn, "Back").Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.Button(th.Theme, &a.saveBtn, "Save").Layout(gtx)
				}),
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				theme := *th.Theme
				theme.TextSize = 20
				return material.Editor(&theme, &a.title, "Title").Layout(gtx)
			})
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.Editor(th.Theme, &a.body, "Write your note...").Layout(gtx)
			})
		}),
	)

	// material.Editor drains input, so we gotta read the text back
	if save {
		a.commit()
	}
	return dims
}

// openNote loads a note into the editors and switches to the edit page.
func (a *App) openNote(i int) {
	a.currentNote = uint(i)
	a.title.SetText(a.Notes[i].Title)
	a.body.SetText(a.Notes[i].Content)
	a.screen = screenEdit
}

// creates an empty note and opens it.
func (a *App) newNote() {
	a.Notes = append(a.Notes, Note{})
	a.openNote(len(a.Notes) - 1)
}

func (a *App) commit() {
	if int(a.currentNote) >= len(a.Notes) {
		return
	}

	n := &a.Notes[a.currentNote]
	n.Title = a.title.Text()
	n.Content = a.body.Text()

	now := time.Now()
	if n.CreatedAt.IsZero() {
		n.CreatedAt = now
	}
	n.ModifiedAt = now

	if a.store.path != "" {
		if err := a.store.save(a.Notes); err != nil {
			log.Printf("notes: save: %v", err)
		}
	}
	a.screen = screenList
}
