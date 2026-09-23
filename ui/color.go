package ui

import "github.com/pterm/pterm"

type Color int

const (
	DEFAULT Color = iota
	RED
	GREEN
	YELLOW
)

var colors = map[Color]pterm.Color{
	DEFAULT: pterm.FgDefault,
	RED:     pterm.FgRed,
	GREEN:   pterm.FgGreen,
	YELLOW:  pterm.FgYellow,
}

func (s Style) PtermColor() pterm.Color {
	return colors[s.Color]
}
