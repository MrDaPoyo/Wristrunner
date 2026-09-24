package notes

import (
	"lib"
	"log"
	"time"
	"uikit"

	"gioui.org/layout"
	"gioui.org/widget"
)

type screen int

const (
	screenList = iota
	screenEdit
)

type Note struct {
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

type App struct {
	Notes []Note

	currentNote uint // index
	screen      screen
	store       store

	list layout.List
	rows []widget.Clickable

	// edit page stuff
	title   widget.Editor
	body    widget.Editor
	backBtn widget.Clickable
	saveBtn widget.Clickable
	addBtn  widget.Clickable
}

func New() uikit.App {
	a := &App{list: layout.List{Axis: layout.Vertical}}
	a.title.SingleLine = true // dont wrap

	dir, err := lib.DataDir("notes")
	if err != nil {
		log.Printf("notes: data dir: %v", err)
		return a // no store: the app still runs, just in memory
	}

	a.store = newStore(dir)
	notes, err := a.store.load()
	if err != nil {
		log.Printf("notes: load: %v", err)
	}
	a.Notes = notes

	return a
}

func (a *App) Layout(gtx layout.Context, th *uikit.Theme) layout.Dimensions {
	if a.screen == screenEdit {
		return a.layoutEdit(gtx, th)
	}
	return a.layoutList(gtx, th)
}

// save on closing because notes would be useless otherwise :P
func (a *App) Close() {
	if a.store.path == "" {
		return
	}
	if err := a.store.save(a.Notes); err != nil {
		log.Printf("notes: save: %v", err)
	}
}
