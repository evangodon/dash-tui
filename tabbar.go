package main

// todo tab bar
// todo tab layout
// todo module

import (
	"strings"

	lg "github.com/charmbracelet/lipgloss"
	"github.com/evangodon/dash/config"
	"github.com/evangodon/dash/ui"
	"github.com/evangodon/dash/util"
)

var (
	tabContainer = lg.NewStyle()
	tabGap       = tabStyle.Copy().
			BorderTop(false).
			BorderLeft(false).
			BorderRight(false)
)

func (m model) BuildTabs(activeTab int, tabs ...config.Tab) string {
	tabboxes := []string{}
	for index, tab := range tabs {
		if index == activeTab {
			tabboxes = append(tabboxes, activeTabStyle.Render(tab.Name))
			continue
		}
		tabboxes = append(tabboxes, tabStyle.Render(tab.Name))
	}

	row := lg.JoinHorizontal(
		lg.Top,
		tabboxes...,
	)

	gap := tabGap.Render(strings.Repeat(" ", util.Max(0, m.window.width-lg.Width(row)-2)))

	row = lg.JoinHorizontal(
		lg.Bottom,
		row,
		gap,
	)

	container := tabContainer.Width(m.window.width)

	return container.Render(row)
}

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
	BorderForeground(ui.Color.Primary)

var activeTabStyle = tabStyle.Copy().
	Border(activeTabBorder, true).
	Bold(true)
