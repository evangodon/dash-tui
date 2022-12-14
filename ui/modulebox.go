package ui

import (
	lg "github.com/charmbracelet/lipgloss"
	"github.com/evangodon/dash/config"
)

var boxStyle = lg.NewStyle().
	Border(lg.RoundedBorder(), true).
	BorderForeground(color.border).
	Padding(0, 1).
	Render

func (cb ComponentBuilder) NewModuleBox(mod config.Module, height int) string {
	b := newBoxWithTitle()

	return b.Render(mod.GetTitle(), mod.Output.String(), mod.GetRenderedWidth(), height)
}
