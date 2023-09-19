# irodori(彩)

<br>
<p align="center">
<img src="img/irodori.png" alt="Color Pallete" height="350" width="350"/>
</p>

<p align="center">
irodori(彩)
</p>

<p align="center">
<a href="https://opensource.org/licenses/MIT">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="MIT License badge">
</a>
<a href="https://github.com/orangekame3/irodori/actions/workflows/tagpr.yml">
<img src="https://github.com/orangekame3/irodori/actions/workflows/tagpr.yml/badge.svg" alt="Tag PR workflow status badge">
</a>
</p>

## Color Pallete

```go
package main

import (
    "fmt"

    "github.com/charmbracelet/lipgloss"
    irodori "github.com/orangekame3/irodori"
    )

    func main() {
    for key, theme := range irodori.Pallete {
        styles, bgStyles := irodori.ColorSample(theme)
        fmt.Println(key)
        fmt.Println(lipgloss.JoinHorizontal(lipgloss.Top, styles...))
        fmt.Println(lipgloss.JoinHorizontal(lipgloss.Top, bgStyles...))
    }
}
```

## Usage

## License

`irodori` is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

## Acknowledgments
