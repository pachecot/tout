// Package tout provides utilities for terminal output.
//
// The tout package includes various functionalities to enhance terminal output,
// such as text coloring using ANSI color codes. It defines the Term type, which
// can be used to manage terminal-related operations and settings.
//
// The package includes a default Cursor that is used by helper functions to
// set the text colors and other display attributes.
//
// example:
//
//	package main
//
//	import (
//		"github.com/pachecot/tout"
//	)
//
//	func main() {
//		tout.Println("Hello, World!")
//	}
//
// The example above prints "Hello, World!" to the terminal using the default
// Cursor settings.
//
// The tout package also provides functions to set the text color, background
// color, and font style of the text. These functions can be used to customize
// the appearance of the text output.
//
// example:
//
//	package main
//
//	import (
//		"github.com/pachecot/tout"
//		"github.com/pachecot/tout/color24"
//	)
//
//	func main() {
//		tout.Foreground(color24.Blue)
//		tout.Println("This is blue text")
//	}
//
// The example above sets the text color to blue and prints "This is blue text"
// to the terminal using the default Cursor settings.
//
// The tout package also provides functions to set the font style of the text.
// These functions can be used to apply bold, italic, underline, and other font
// styles to the text output.
//
// example:
//
//	package main
//
//	import (
//		"github.com/pachecot/tout"
//	)
//
//	func main() {
//		tout.SetFont(tout.BoldFont | tout.UnderlineFont)
//		tout.Println("This is bold underlined text")
//	}
//
// The example above sets the font style to bold and underline and prints "This
// is bold underlined text" to the terminal using the default Cursor settings.
package tout
