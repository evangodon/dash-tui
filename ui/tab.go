package ui

import lg "github.com/charmbracelet/lipgloss"

var (
	activeTabBorder = lg.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}

	tabBorder = lg.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}
)

var tabBottomBorder = lg.Border{
	Bottom: "🬂",
}

var tabStyle = lg.NewStyle().
	Padding(0, 1).
	Border(tabBorder, true).
	BorderForeground(color.primary)

var activeTabStyle = tabStyle.Copy().
	Border(activeTabBorder, true).
	Bold(true)
