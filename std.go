package tout

import (
	"io"
	"os"

	"github.com/pachecot/tout/color24"
)

var (
	out *Cursor
	err *Cursor
)

func init() {
	out = NewCursor(os.Stdout)
	err = NewCursor(os.Stderr)
	err.Foreground(color24.Red)
}

// DefaultCursor returns sets the default cursor for terminal output.
func DefaultCursor() *Cursor {
	return out
}

// SetDefaultCursor sets the default cursor for terminal output.
func SetDefaultCursor(c *Cursor) {
	out = c
}

// Writer returns the writer associated with the default cursor.
func Writer() io.Writer {
	return out.Writer()
}

// ErrorCursor returns the cursor used for error output.
func ErrorCursor() *Cursor {
	return err
}

// SetErrorCursor sets the cursor for error output.
func SetErrorCursor(c *Cursor) {
	err = c
}

// ErrorWriter returns the writer associated with the error cursor.
func ErrorWriter() io.Writer {
	return err.Writer()
}

// Font returns the current font type of the default cursor.
func Font() FontType {
	return out.Font
}

// SetFont sets the font type for the default cursor.
func SetFont(f FontType) {
	out.Font = f
}

// ResetFont resets the font type of the default cursor to the default value.
func ResetFont() {
	out.Font = 0
}

// AddStyle adds a style to the current font type of the default cursor.
func AddStyle(f FontType) {
	out.Font |= f
}

// RemoveStyle removes a style from the current font type of the default cursor.
func RemoveStyle(f FontType) {
	out.Font &= ^f
}

// Foreground sets the foreground color of the default cursor.
func Foreground(c Color) {
	out.Foreground(c)
}

// ResetForeground resets the foreground color of the default cursor to the default value.
func ResetForeground() {
	out.ResetForeground()
}

// DefaultForeground sets the foreground color of the default cursor to the default value.
func DefaultForeground() {
	out.DefaultForeground()
}

// Background sets the background color of the default cursor.
func Background(c Color) {
	out.Background(c)
}

// ResetBackground resets the background color of the default cursor to the default value.
func ResetBackground() {
	out.ResetBackground()
}

// DefaultBackground sets the background color of the default cursor to the default value.
func DefaultBackground() {
	out.DefaultBackground()
}

// Write writes the given string or byte slice to the default cursor's writer.
func Write[T string | []byte](t T) {
	out.w.Write([]byte(t))
}

// Print prints the given arguments to the default cursor's writer.
func Print(a ...any) {
	out.Print(a...)
}

// Println prints the given arguments to the default cursor's writer with a newline.
func Println(a ...any) {
	out.Println(a...)
}

// Printf formats according to a format specifier and writes to the default cursor's writer.
func Printf(format string, a ...any) {
	out.Printf(format, a...)
}

// Error prints the given arguments to the error cursor's writer.
func Error(a ...any) {
	err.Print(a...)
}

// Errorln prints the given arguments to the error cursor's writer with a newline.
func Errorln(a ...any) {
	err.Println(a...)
}

// Errorf formats according to a format specifier and writes to the error cursor's writer.
func Errorf(format string, a ...any) {
	err.Printf(format, a...)
}
