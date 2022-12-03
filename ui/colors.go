package ui

import lg "github.com/charmbracelet/lipgloss"

type colors struct {
	primary lg.AdaptiveColor
	border  lg.AdaptiveColor
}

var color = colors{
	primary: lg.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"},
	border:  lg.AdaptiveColor{Light: "#bcc0cc", Dark: "#45475a"},
}
