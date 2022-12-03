package ui

import (
	"strings"

	lg "github.com/charmbracelet/lipgloss"
)

var tabStyle = lg.NewStyle().
	Border(lg.NormalBorder(), true).
	BorderForeground(color.primary).
	Padding(0, 1)

var activeTabStyle = tabStyle.Copy().Border(lg.DoubleBorder(), true)

func (u UI) BuildTabs(activeTab int, tabs ...string) string {
	doc := strings.Builder{}

	rendered := []string{}
	for index, tab := range tabs {
		if index == activeTab {
			rendered = append(rendered, activeTabStyle.Render(tab))
			continue
		}
		rendered = append(rendered, tabStyle.Render(tab))
	}

	row := lg.JoinHorizontal(
		lg.Top,
		rendered...,
	)

	doc.WriteString(row + "\n\n")
	return doc.String()
}
