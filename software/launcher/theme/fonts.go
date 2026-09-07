package theme

import (
	"io/fs"

	"gioui.org/font"
	"gioui.org/font/opentype"

	"wr-ui/assets"
)

// loadFont parses a single font file into a face.
func loadFont(data []byte) []font.FontFace {
	face, err := opentype.Parse(data)
	if err != nil {
		panic(err)
	}
	return []font.FontFace{{Font: font.Font{Typeface: face.Font().Typeface}, Face: face}}
}

// LoadFonts returns the font faces bundled into the binary.
func LoadFonts() (fonts []font.FontFace, err error) {
	entries, err := fs.Glob(assets.Fonts, "fonts/*.ttf")
	if err != nil {
		return nil, err
	}

	for _, name := range entries {
		data, err := assets.Fonts.ReadFile(name)
		if err != nil {
			return nil, err
		}
		fonts = append(fonts, loadFont(data)...)
	}
	return fonts, nil
}
