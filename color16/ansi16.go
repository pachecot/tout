// Package color16 provides a set of basic 16-color ANSI color definitions.
//
// The package color16 defines a set of basic ANSI colors that can be used
// for terminal text coloring. Each color is represented by a constant of
// type Color, which is an integer.
package color16

type Color int

const (
	Black   Color = 0  // #000000
	Maroon  Color = 1  // #800000
	Green   Color = 2  // #008000
	Olive   Color = 3  // #808000
	Navy    Color = 4  // #000080
	Purple  Color = 5  // #800080
	Teal    Color = 6  // #008080
	Silver  Color = 7  // #c0c0c0
	Grey    Color = 8  // #808080
	Red     Color = 9  // #ff0000
	Lime    Color = 10 // #00ff00
	Yellow  Color = 11 // #ffff00
	Blue    Color = 12 // #0000ff
	Fuchsia Color = 13 // #ff00ff
	Aqua    Color = 14 // #00ffff
	White   Color = 15 // #ffffff

	// alias color names

	Salami           Color = 1
	GreenHills       Color = 2
	MothGreen        Color = 2
	MongolianPlateau Color = 3
	SwampGreen       Color = 3
	VerdeTropa       Color = 3
	DeepBlue         Color = 4
	Ultraberry       Color = 5
	BellyFlop        Color = 6
	Macquarie        Color = 6
	PondBath         Color = 6
	Windows95Desktop Color = 6
	NeoTokyoGrey     Color = 7
	SilverTippedSage Color = 7
	StonewallGrey    Color = 7
	Waxwing          Color = 7
)

const (
	fgDefault = "39"
	bgDefault = "49"

	fgBlack         = "30"
	fgRed           = "31"
	fgGreen         = "32"
	fgBrown         = "33"
	fgBlue          = "34"
	fgMagenta       = "35"
	fgCyan          = "36"
	fgWhite         = "37"
	fgBrightBlack   = "90"
	fgBrightRed     = "91"
	fgBrightGreen   = "92"
	fgBrightBrown   = "93"
	fgBrightBlue    = "94"
	fgBrightMagenta = "95"
	fgBrightCyan    = "96"
	fgBrightWhite   = "97"

	bgBlack         = "40"
	bgRed           = "41"
	bgGreen         = "42"
	bgBrown         = "43"
	bgBlue          = "44"
	bgMagenta       = "45"
	bgCyan          = "46"
	bgWhite         = "47"
	bgBrightBlack   = "100"
	bgBrightRed     = "101"
	bgBrightGreen   = "102"
	bgBrightBrown   = "103"
	bgBrightBlue    = "104"
	bgBrightMagenta = "105"
	bgBrightCyan    = "106"
	bgBrightWhite   = "107"
)

var (
	fgColors = []string{
		fgBlack,
		fgRed,
		fgGreen,
		fgBrown,
		fgBlue,
		fgMagenta,
		fgCyan,
		fgWhite,
		fgBrightBlack,
		fgBrightRed,
		fgBrightGreen,
		fgBrightBrown,
		fgBrightBlue,
		fgBrightMagenta,
		fgBrightCyan,
		fgBrightWhite,
	}

	bgColors = []string{
		bgBlack,
		bgRed,
		bgGreen,
		bgBrown,
		bgBlue,
		bgMagenta,
		bgCyan,
		bgWhite,
		bgBrightBlack,
		bgBrightRed,
		bgBrightGreen,
		bgBrightBrown,
		bgBrightBlue,
		bgBrightMagenta,
		bgBrightCyan,
		bgBrightWhite,
	}
)

// Foreground returns the ANSI escape code for setting the foreground color
// to the Color value. The color is specified using ANSIs Basic 16 colors.
// If the color is not valid, it returns the default foreground color.
func (c Color) Foreground() string {
	if c < 0 || int(c) >= len(fgColors) {
		return fgDefault
	}
	return fgColors[c]
}

// Background returns the ANSI escape code for setting the background color
// to the Color value. The color is specified using ANSIs Basic 16 colors.
// If the color is not valid, it returns the default background color.
func (c Color) Background() string {
	if c < 0 || int(c) >= len(bgColors) {
		return bgDefault
	}
	return bgColors[c]
}
