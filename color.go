package tout

import (
	"github.com/jwalton/go-supportscolor"
	"github.com/pachecot/tout/color16"
	"github.com/pachecot/tout/color24"
	"github.com/pachecot/tout/color256"
)

// Color interface represents a color that can provide ANSI escape codes for foreground and background colors.
type Color interface {
	Foreground() string
	Background() string
}

// colorLevel determines the color support level for a given Color.
func colorLevel(k Color) supportscolor.ColorLevel {
	switch k.(type) {
	case color16.Color:
		return supportscolor.Basic
	case color256.Color:
		return supportscolor.Ansi256
	case color24.Color:
		return supportscolor.Ansi16m
	}
	return supportscolor.None
}
