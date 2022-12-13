package ui

import (
	"strings"

	lg "github.com/charmbracelet/lipgloss"
)

var (
	tabContainer = lg.NewStyle()
	tabGap       = tabStyle.Copy().
			BorderTop(false).
			BorderLeft(false).
			BorderRight(false)
)

func (cb ComponentBuilder) BuildTabs(activeTab int, tabs ...string) string {
	tabboxes := []string{}
	for index, tab := range tabs {
		if index == activeTab {
			tabboxes = append(tabboxes, activeTabStyle.Render(tab))
			continue
		}
		tabboxes = append(tabboxes, tabStyle.Render(tab))
	}

	row := lg.JoinHorizontal(
		lg.Top,
		tabboxes...,
	)

	gap := tabGap.Render(strings.Repeat(" ", Max(0, cb.window.Width-lg.Width(row)-2)))

	row = lg.JoinHorizontal(
		lg.Bottom,
		row,
		gap,
	)

	container := tabContainer.Width(cb.window.Width)

	return container.Render(row)
}
