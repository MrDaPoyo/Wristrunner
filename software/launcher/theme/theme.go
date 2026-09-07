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
