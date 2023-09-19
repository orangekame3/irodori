// Package irodori is a color library for Go (golang).
package irodori

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// Pallete is a map of color themes.
var Pallete = map[string]Theme{
	// Japanese color themes
	"Zen":       Zen,
	"Akebono":   Akebono,
	"Shinonome": Shinonome,
	"Sakura":    Sakura,
	"Yugure":    Yugure,
	"Yamabuki":  Yamabuki,
	// general color themes
	"Nord":        Nord,
	"GitHubDark":  GitHubDark,
	"GitHubLight": GitHubLight,
	"Monokai":     Monokai,
	"Pastel":      Pastel,
	"Calm":        Calm,
}

// Theme is a color theme.
type Theme struct {
	Background         Color
	PrimaryText        Color
	SecondaryText      Color
	PrimaryHighlight   Color
	SecondaryHighlight Color
}

// Color is a color.
type Color struct {
	Hex string
	RGB [3]int
}

// Nord color theme
var Nord = Theme{
	Background:         Color{Hex: "#2E3440", RGB: [3]int{46, 52, 64}},
	PrimaryText:        Color{Hex: "#ECEFF4", RGB: [3]int{236, 239, 244}},
	SecondaryText:      Color{Hex: "#5E81AC", RGB: [3]int{94, 129, 172}},
	PrimaryHighlight:   Color{Hex: "#A3BE8C", RGB: [3]int{63, 190, 140}},
	SecondaryHighlight: Color{Hex: "#B48EAD", RGB: [3]int{180, 142, 173}},
}

// GitHubDark color theme
var GitHubDark = Theme{
	Background:         Color{Hex: "#0D1117", RGB: [3]int{13, 17, 23}},
	PrimaryText:        Color{Hex: "#C9D1D9", RGB: [3]int{201, 209, 217}},
	SecondaryText:      Color{Hex: "#8B949E", RGB: [3]int{139, 148, 158}},
	PrimaryHighlight:   Color{Hex: "#58A6FF", RGB: [3]int{88, 166, 255}},
	SecondaryHighlight: Color{Hex: "#F9826C", RGB: [3]int{249, 130, 108}},
}

// GitHubLight color theme
var GitHubLight = Theme{
	Background:         Color{Hex: "#FFFFFF", RGB: [3]int{255, 255, 255}},
	PrimaryText:        Color{Hex: "#24292E", RGB: [3]int{36, 41, 46}},
	SecondaryText:      Color{Hex: "#586069", RGB: [3]int{88, 96, 105}},
	PrimaryHighlight:   Color{Hex: "#E36209", RGB: [3]int{227, 98, 9}},
	SecondaryHighlight: Color{Hex: "#0366D6", RGB: [3]int{3, 102, 214}},
}

// Monokai color theme
var Monokai = Theme{
	Background:         Color{Hex: "#272822", RGB: [3]int{39, 40, 34}},
	PrimaryText:        Color{Hex: "#F8F8F2", RGB: [3]int{248, 248, 242}},
	SecondaryText:      Color{Hex: "#75715E", RGB: [3]int{117, 113, 94}},
	PrimaryHighlight:   Color{Hex: "#F92672", RGB: [3]int{249, 38, 114}},
	SecondaryHighlight: Color{Hex: "#A6E22E", RGB: [3]int{166, 226, 46}},
}

// Pastel color theme
var Pastel = Theme{
	Background:         Color{Hex: "#FDF6E3", RGB: [3]int{253, 246, 227}},
	PrimaryText:        Color{Hex: "#657B83", RGB: [3]int{101, 123, 131}},
	SecondaryText:      Color{Hex: "#93A1A1", RGB: [3]int{147, 161, 161}},
	PrimaryHighlight:   Color{Hex: "#B58900", RGB: [3]int{181, 137, 0}},
	SecondaryHighlight: Color{Hex: "#2AA198", RGB: [3]int{42, 161, 152}},
}

// Calm color theme
var Calm = Theme{
	Background:         Color{Hex: "#F5F5F5", RGB: [3]int{245, 245, 245}},
	PrimaryText:        Color{Hex: "#2E4057", RGB: [3]int{46, 64, 87}},
	SecondaryText:      Color{Hex: "#4A6A8F", RGB: [3]int{74, 106, 143}},
	PrimaryHighlight:   Color{Hex: "#81A1C1", RGB: [3]int{129, 161, 193}},
	SecondaryHighlight: Color{Hex: "#A0C8E2", RGB: [3]int{160, 200, 226}},
}

// Zen color theme
var Zen = Theme{
	Background:         Color{Hex: "#D3CFC7", RGB: [3]int{211, 207, 199}},
	PrimaryText:        Color{Hex: "#59524C", RGB: [3]int{89, 82, 76}},
	SecondaryText:      Color{Hex: "#7B6D63", RGB: [3]int{123, 109, 99}},
	PrimaryHighlight:   Color{Hex: "#A89984", RGB: [3]int{168, 153, 132}},
	SecondaryHighlight: Color{Hex: "#D5C4A1", RGB: [3]int{213, 196, 161}},
}

// Akebono color theme
var Akebono = Theme{
	Background:         Color{Hex: "#F5F5F5", RGB: [3]int{245, 245, 245}},
	PrimaryText:        Color{Hex: "#373B41", RGB: [3]int{55, 59, 65}},
	SecondaryText:      Color{Hex: "#969896", RGB: [3]int{150, 152, 150}},
	PrimaryHighlight:   Color{Hex: "#CC6666", RGB: [3]int{204, 102, 102}},
	SecondaryHighlight: Color{Hex: "#8ABEB7", RGB: [3]int{138, 190, 183}},
}

// Shinonome color theme
var Shinonome = Theme{
	Background:         Color{Hex: "#FFFFFF", RGB: [3]int{255, 255, 255}},
	PrimaryText:        Color{Hex: "#34434B", RGB: [3]int{52, 67, 75}},
	SecondaryText:      Color{Hex: "#7A8B8F", RGB: [3]int{122, 139, 143}},
	PrimaryHighlight:   Color{Hex: "#3E8D96", RGB: [3]int{62, 141, 150}},
	SecondaryHighlight: Color{Hex: "#A09F91", RGB: [3]int{160, 159, 145}},
}

// Sakura color theme
var Sakura = Theme{
	Background:         Color{Hex: "#F5F5F5", RGB: [3]int{245, 245, 245}},
	PrimaryText:        Color{Hex: "#333333", RGB: [3]int{51, 51, 51}},
	SecondaryText:      Color{Hex: "#777777", RGB: [3]int{119, 119, 119}},
	PrimaryHighlight:   Color{Hex: "#D6899A", RGB: [3]int{214, 137, 154}},
	SecondaryHighlight: Color{Hex: "#F2D1D2", RGB: [3]int{242, 209, 210}},
}

// Yugure color theme
var Yugure = Theme{
	Background:         Color{Hex: "#2E2633", RGB: [3]int{46, 38, 51}},
	PrimaryText:        Color{Hex: "#ABA9B2", RGB: [3]int{171, 169, 178}},
	SecondaryText:      Color{Hex: "#CBC9D0", RGB: [3]int{203, 201, 208}},
	PrimaryHighlight:   Color{Hex: "#D38342", RGB: [3]int{211, 131, 66}},
	SecondaryHighlight: Color{Hex: "#B9AEB2", RGB: [3]int{185, 174, 178}},
}

// Yamabuki color theme
var Yamabuki = Theme{
	Background:         Color{Hex: "#FAF3DC", RGB: [3]int{250, 243, 220}},
	PrimaryText:        Color{Hex: "#8B4513", RGB: [3]int{139, 69, 19}},
	SecondaryText:      Color{Hex: "#AA8C6F", RGB: [3]int{170, 140, 111}},
	PrimaryHighlight:   Color{Hex: "#DAA520", RGB: [3]int{218, 165, 32}},
	SecondaryHighlight: Color{Hex: "#9932CC", RGB: [3]int{153, 50, 204}},
}

// GetHex returns the hex code of the color.
func (c Color) GetHex() string {
	return c.Hex
}

// GetRGB returns the RGB code of the color.
func (c Color) GetRGB() [3]int {
	return c.RGB
}

// Role returns the color of the given role.
func (t Theme) Role(r string) (Color, error) {
	switch r {
	case "Background":
		return t.Background, nil
	case "PrimaryText":
		return t.PrimaryText, nil
	case "SecondaryText":
		return t.SecondaryText, nil
	case "PrimaryHighlight":
		return t.PrimaryHighlight, nil
	case "SecondaryHighlight":
		return t.SecondaryHighlight, nil
	default:
		return Color{}, fmt.Errorf("unknown role: %s", r)
	}
}

// ColorSample returns color samples for the given theme.
func ColorSample(theme Theme) (styles []string, bgStyles []string) {
	roles := []string{"PrimaryText", "SecondaryText", "PrimaryHighlight", "SecondaryHighlight"}
	bgColor, _ := theme.Role("Background")
	for _, r := range roles {
		textColor, _ := theme.Role(r)
		styles = append(styles, format(bgColor, textColor, r))
		bgStyles = append(bgStyles, format(textColor, textColor, ""))
	}
	return styles, bgStyles
}

// format returns a string formatted with the given background and text colors.
func format(bgColor, textColor Color, r string) string {
	return lipgloss.NewStyle().
		Background(lipgloss.Color(bgColor.GetHex())).
		Foreground(lipgloss.Color(textColor.GetHex())).
		AlignHorizontal(lipgloss.Center).
		AlignVertical(lipgloss.Center).
		Padding(0, 1).
		Height(1).
		Width(25).
		Render(r)
}
