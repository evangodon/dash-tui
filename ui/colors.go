package ui

import lg "github.com/charmbracelet/lipgloss"

type colors struct {
	primary lg.AdaptiveColor
}

var color = colors{
	primary: lg.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"},
}
