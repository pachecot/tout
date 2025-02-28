package tout

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/jwalton/go-supportscolor"
)

// FontType represents different font styles that can be applied to terminal text.
type FontType int

// Font style constants.
const (
	BoldFont      FontType = 1 << iota // Bold font style
	LightFont                          // Light font style
	ItalicFont                         // Italic font style
	UnderlineFont                      // Underline font style
	BlinkFont                          // Blink font style
	ReversedFont                       // Reversed font style
	CrossOutFont                       // Crossed-out font style
)

// Has checks if the FontType includes a specific style.
func (f FontType) Has(s FontType) bool {
	return f&s == s
}

// ANSI escape codes and sequences.
const (
	ESC       = "\x1b"                   // Escape character
	CSI       = ESC + "["                // Control Sequence Introducer
	csiReset  = ESC + "[0m"              // Reset all attributes
	csiFormat = CSI + "%sm%s" + csiReset // Format string for applying styles

	csiBold       = "1" // Set bold
	csiLight      = "2" // Set half-bright
	csiItalic     = "3" // Set italic
	csiUnderscore = "4" // Set underscore
	csiBlink      = "5" // Set blink
	csiReverse    = "7" // Set reverse video
	csiCrossOut   = "9" // Set strikeout / cross-out

	csiDefaultFg = "39" // set default Fg
	csiDefaultBg = "49" // set default Bg
)

// Cursor represents a terminal cursor with associated writer, colors, and font styles.
type Cursor struct {
	w     io.Writer
	fg    string
	bg    string
	Font  FontType
	Level supportscolor.ColorLevel
}

// NewCursor creates a new Cursor with the specified writer.
// It detects the color support level of the writer if it is a file.
//
// If the writer is nil, it uses os.Stdout as the default writer.
func NewCursor(w io.Writer) *Cursor {
	if w == nil {
		w = os.Stdout
	}
	if f, ok := w.(*os.File); ok {
		support := supportscolor.SupportsColor(f.Fd())
		return &Cursor{
			w:     w,
			Level: support.Level,
		}
	}
	return &Cursor{
		w:     w,
		Level: supportscolor.None,
	}
}

// SetWriter sets the writer for the Cursor and updates the color support level.
func (c *Cursor) SetWriter(w io.Writer) {
	c.w = w
	c.Level = supportscolor.None
	if f, ok := w.(*os.File); ok {
		support := supportscolor.SupportsColor(f.Fd())
		c.Level = support.Level
	}
}

// Writer returns the writer associated with the Cursor.
func (c *Cursor) Writer() io.Writer {
	return c.w
}

// Foreground sets the foreground color of the Cursor.
// If the color level is not supported, it resets the foreground color.
func (c *Cursor) Foreground(k Color) {
	if c.Level < colorLevel(k) {
		c.fg = ""
		return
	}
	c.fg = k.Foreground()
}

// Background sets the background color of the Cursor.
// If the color level is not supported, it resets the background color.
func (c *Cursor) Background(k Color) {
	if c.Level < colorLevel(k) {
		c.bg = ""
		return
	}
	c.bg = k.Background()
}

// DefaultForeground sets the foreground color of the Cursor to the default value.
func (c *Cursor) DefaultForeground() {
	c.fg = csiDefaultFg
}

// DefaultBackground sets the background color of the Cursor to the default value.
func (c *Cursor) DefaultBackground() {
	c.bg = csiDefaultBg
}

// ResetForeground resets the foreground color of the Cursor.
func (c *Cursor) ResetForeground() {
	c.fg = ""
}

// ResetBackground resets the background color of the Cursor.
func (c *Cursor) ResetBackground() {
	c.bg = ""
}

func (c *Cursor) params() []byte {
	if c.Level == supportscolor.None {
		return make([]byte, 0)
	}
	var b bytes.Buffer
	if len(c.fg) > 0 {
		b.WriteString(c.fg)
		b.WriteByte(';')
	}
	if len(c.bg) > 0 {
		b.WriteString(c.bg)
		b.WriteByte(';')
	}
	flags := c.Font
	for flags > 0 {
		next := flags & (flags - 1)
		switch flags ^ next {
		case BlinkFont:
			b.WriteString(csiBlink)
		case BoldFont:
			b.WriteString(csiBold)
		case CrossOutFont:
			b.WriteString(csiCrossOut)
		case UnderlineFont:
			b.WriteString(csiUnderscore)
		case ItalicFont:
			b.WriteString(csiItalic)
		case LightFont:
			b.WriteString(csiLight)
		case ReversedFont:
			b.WriteString(csiReverse)
		}
		b.WriteByte(';')
		flags = next
	}
	if b.Len() == 0 {
		return make([]byte, 0)
	}
	p := b.Bytes()
	return p[:len(p)-1]
}

// render applies the current styles and colors to the given byte slice.
func (c *Cursor) render(p []byte) []byte {
	ps := c.params()
	if len(ps) == 0 {
		return p
	}
	return fmt.Appendf(nil, csiFormat, ps, p)
}

// Print prints the given arguments to the Cursor's writer.
func (c *Cursor) Print(a ...any) {
	p := fmt.Append(nil, a...)
	c.w.Write(c.render(p))
}

// Println prints the given arguments to the Cursor's writer with a newline.
// The newline is not styled.
func (c *Cursor) Println(a ...any) {
	p := fmt.Append(nil, a...)
	c.w.Write(c.render(p))
	c.w.Write([]byte{'\n'})
}

// Printf formats according to a format specifier and writes to the Cursor's writer.
func (c *Cursor) Printf(format string, a ...any) {
	p := fmt.Appendf(nil, format, a...)
	c.w.Write(c.render(p))
}
