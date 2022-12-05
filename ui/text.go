package ui

import (
	lg "github.com/charmbracelet/lipgloss"
)

var Italic = lg.NewStyle().Italic(true)

var AppTitle = Italic.Faint(true).Render
