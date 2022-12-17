package ui

import lg "github.com/charmbracelet/lipgloss"

type colors struct {
	primary lg.AdaptiveColor
	border  lg.AdaptiveColor
	green   lg.AdaptiveColor
	yellow  lg.AdaptiveColor
	red     lg.AdaptiveColor
}

var color = colors{
	primary: lg.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"},
	border:  lg.AdaptiveColor{Light: "#bcc0cc", Dark: "#a6adc8"},
	green:   lg.AdaptiveColor{Light: "#a6e3a1", Dark: "#a6e3a1"},
	yellow:  lg.AdaptiveColor{Light: "#f9e2af", Dark: "#f9e2af"},
	red:     lg.AdaptiveColor{Light: "#f38ba8", Dark: "#f38ba8"},
}

var (
	ColorError = color.red
)
