package ui

import (
	"strings"

	lg "github.com/charmbracelet/lipgloss"
)

var tabStyle = lg.NewStyle().
	Padding(0, 1)

var activeTabStyle = tabStyle.Copy().
	Background(color.primary)

var tabBottomBorder = lg.Border{
	Bottom: "🬂",
}

var tabbarStyle = lg.NewStyle().
	Border(tabBottomBorder, false, false, true).
	BorderForeground(color.primary)

func (cb ComponentBuilder) BuildTabs(activeTab int, tabs ...string) string {
	doc := strings.Builder{}

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
		strings.Join(tabboxes, " • "),
	)

	out := tabbarStyle.Width(cb.window.Width).Render(row)

	doc.WriteString(out)
	doc.WriteString("\n\n")
	return doc.String()
}
