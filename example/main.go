// pakcage main is example of irodori
package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/orangekame3/irodori"
)

func main() {
	for key, theme := range irodori.Pallete {
		styles, bgStyles := irodori.ColorSample(theme)
		fmt.Println(key)
		fmt.Println(lipgloss.JoinHorizontal(lipgloss.Top, styles...))
		fmt.Println(lipgloss.JoinHorizontal(lipgloss.Top, bgStyles...))
	}
}
