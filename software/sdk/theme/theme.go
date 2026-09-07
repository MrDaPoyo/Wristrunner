package theme

import (
	"image/color"

	"gioui.org/text"
	"gioui.org/widget/material"
)

func Vectorheart() *material.Theme {
	th := material.NewTheme() // base

	fonts, err := LoadFonts()
	if err != nil {
		panic(err)
	}

	th.Shaper = text.NewShaper(text.WithCollection(fonts))
	th.Face = "Oxanium"

	// override the color palette
	th.Palette = material.Palette{
		Bg:         rgb(0x000000), // background
		Fg:         rgb(0xFFFFFF), // text/foreground
		ContrastBg: rgb(0x6750A4), // accent (buttons, etc.)
		ContrastFg: rgb(0xFFFFFF),
	}

	// default text size
	th.TextSize = 16

	return th
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
