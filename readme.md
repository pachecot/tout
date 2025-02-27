# tout

`tout` is a Go package that provides utilities for styled terminal output.

## Features

- **Text Coloring**: Apply ANSI color codes to text for enhanced readability and visual appeal.
- **Cursor Management**: Manage terminal cursor settings and operations.
- **Font Styling**: Set and modify font styles for terminal text.
- **Error Handling**: Separate cursors and writers for standard and error output.

## Installation

To install the `tout` package, use the following command:

```sh
go get github.com/pachecot/tout
```

## Usage

### Basic Example

```go
package main

import (
    "github.com/pachecot/tout"
    "github.com/pachecot/tout/color16"
)

func main() {
    tout.Foreground(color16.Red)
    tout.SetFont(tout.BoldFont | tout.UnderlineFont)
    tout.Println("This is bold red underlined text")
    tout.ResetForeground()
	tout.ResetFont()
    tout.Println("This is default text")
}
```

### Setting Colors

```go
import "github.com/pachecot/tout/color24"

tout.Foreground(color24.Magenta)
tout.Println("This is magenta text")

tout.Foreground(color24.RGB(128,128,128))
tout.Println("This is grey text")

tout.Foreground(color24.Hex("#00ffd7"))
tout.Println("This is bright teal text")
```

### Error Output

```go
tout.Errorln("This is an error message")
```

## Documentation

For detailed documentation, please refer to the [GoDoc](https://pkg.go.dev/github.com/pachecot/tout).

## License

This project is licensed under the MIT License. See the LICENSE file for details.
