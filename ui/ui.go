package ui

import lg "github.com/charmbracelet/lipgloss"

type colors struct {
	Primary lg.AdaptiveColor
	Border  lg.AdaptiveColor
	Green   lg.AdaptiveColor
	Yellow  lg.AdaptiveColor
	Red     lg.AdaptiveColor
}

var Color = colors{
	Primary: lg.AdaptiveColor{Light: "#1e66f5	", Dark: "#89b4fa"},
	Border: lg.AdaptiveColor{Light: "#bcc0cc", Dark: "#a6adc8"},
	Green:  lg.AdaptiveColor{Light: "#a6e3a1", Dark: "#a6e3a1"},
	Yellow: lg.AdaptiveColor{Light: "#f9e2af", Dark: "#f9e2af"},
	Red:    lg.AdaptiveColor{Light: "#f38ba8", Dark: "#f38ba8"},
}

var (
	ColorError = Color.Red
)

var Italic = lg.NewStyle().Italic(true)
var BoldText = lg.NewStyle().Bold(true).Render
var AppTitle = Italic.Faint(true).Render

var GreenText = lg.NewStyle().Foreground(Color.Green).Render
var YellowText = lg.NewStyle().Foreground(Color.Yellow).Render
var RedText = lg.NewStyle().Foreground(Color.Red).Render
