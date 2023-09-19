# 🚧【WIP】

<br>
<p align="center">
irodori(彩)
</p>

<p align="center">
<a href="https://opensource.org/licenses/MIT">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="MIT License badge">
</a>
<a href="https://pkg.go.dev/github.com/orangekame3/irodori">
<img src="https://github.com/orangekame3/irodori/actions/workflows/release.yml/badge.svg" alt="Release workflow status badge">
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
    irodori "github.com/oragekame3/irodori"
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

<p align="center">
<img src="img/sample.png" alt="Viewer of Diff" height="500"/>
</p>

## Usage

## License

`irodori` is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

## Acknowledgments
