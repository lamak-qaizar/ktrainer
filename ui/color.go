package ui

import "github.com/pterm/pterm"

type Color int

const (
	DEFAULT Color = iota
	RED
	GREEN
	YELLOW
	GRAY
)

var colors = map[Color]pterm.Color{
	DEFAULT: pterm.FgDefault,
	RED:     pterm.FgRed,
	GREEN:   pterm.FgGreen,
	YELLOW:  pterm.FgYellow,
	GRAY:    pterm.FgGray,
}

func (c Color) toPterm() pterm.Color {
	return colors[c]
}
