package uikit

import (
	"image/color"

	"gioui.org/text"
	"gioui.org/widget/material"
)

type Theme struct {
	*material.Theme
	Bg, Surface, Accent, Text, Muted color.NRGBA
	Spacing                          int // base unit in dp
}

func VectorheartTheme() *Theme {
	th := material.NewTheme() // base

	fonts, err := LoadFonts()
	if err != nil {
		panic(err)
	}

	th.Shaper = text.NewShaper(text.WithCollection(fonts))
	th.Face = "Michroma" // available fonts: Michroma, Oxanium

	// override the color palette
	th.Palette = material.Palette{
		Bg:         rgb(0x171717), // background
		Fg:         rgb(0xdedede), // text/foreground
		ContrastBg: rgb(0xf2572b), // accent (buttons, etc.)
		ContrastFg: rgb(0x4d4d4d),
	}

	// default text size
	th.TextSize = 16

	return &Theme{
		Theme:   th,
		Bg:      color.NRGBA{R: 0x0d, G: 0x11, B: 0x17, A: 0xff}, // dark cyberdeck palette
		Surface: color.NRGBA{R: 0x16, G: 0x1c, B: 0x24, A: 0xff},
		Accent:  color.NRGBA{R: 0x39, G: 0xff, B: 0x14, A: 0xff}, // terminal green
		Text:    color.NRGBA{R: 0xe0, G: 0xe0, B: 0xe0, A: 0xff},
		Muted:   color.NRGBA{R: 0x80, G: 0x88, B: 0x90, A: 0xff},
		Spacing: 8,
	}
}

// converts a hex code to color.NRGBA
func rgb(c uint32) color.NRGBA {
	return color.NRGBA{
		R: uint8(c >> 16),
		G: uint8(c >> 8),
		B: uint8(c),
		A: 0xFF,
	}
}
