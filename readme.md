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
go get github.com/tpacheco/tout
```

## Usage

### Basic Example

```go
package main

import (
    "github.com/tpacheco/tout"
    "github.com/tpacheco/tout/color16"
)

func main() {
    tout.SetDefaultCursor(tout.NewCursor())
    tout.Foreground(color16.Red)
    tout.Println("This is red text")
    tout.ResetForeground()
    tout.Println("This is default text")
}
```

### Setting Colors

```go
import "github.com/tpacheco/tout/color24"

tout.Foreground(color24.Magenta)
tout.Println("This is magenta text")
tout.ResetForeground()
```

### Error Output

```go
tout.SetErrorCursor(tout.NewCursor())
tout.Errorln("This is an error message")
```

## Documentation

For detailed documentation, please refer to the [GoDoc](https://pkg.go.dev/github.com/tpacheco/tout).

## License

This project is licensed under the MIT License. See the LICENSE file for details.
