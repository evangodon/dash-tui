package ui

import (
	lg "github.com/charmbracelet/lipgloss"
)

var Italic = lg.NewStyle().Italic(true)
var BoldText = lg.NewStyle().Bold(true).Render
var AppTitle = Italic.Faint(true).Render

var GreenText = lg.NewStyle().Foreground(color.green).Render
var YellowText = lg.NewStyle().Foreground(color.yellow).Render
var RedText = lg.NewStyle().Foreground(color.red).Render
