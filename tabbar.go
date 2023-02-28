package main

// todo tab bar
// todo tab layout
// todo module

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	lg "github.com/charmbracelet/lipgloss"
	"github.com/evangodon/dash/config"
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
	borderColor := lipgloss.Color(m.config.Settings.PrimaryColor)
	for index, tab := range tabs {
		if index == activeTab {
			tabboxes = append(tabboxes, activeTabStyle.BorderForeground(borderColor).Render(tab.Name))
			continue
		}
		tabboxes = append(tabboxes, tabStyle.BorderForeground(borderColor).Render(tab.Name))
	}

	row := lg.JoinHorizontal(
		lg.Top,
		tabboxes...,
	)

	gap := tabGap.BorderForeground(borderColor).Render(strings.Repeat(" ", util.Max(0, m.window.width-lg.Width(row)-2)))

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
	Border(tabBorder, true)

var activeTabStyle = tabStyle.Copy().
	Border(activeTabBorder, true).
	Bold(true)
