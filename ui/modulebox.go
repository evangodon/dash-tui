package ui

import (
	lg "github.com/charmbracelet/lipgloss"

	"github.com/evangodon/dash/module"
)

var boxStyle = lg.NewStyle().
	Border(lg.RoundedBorder(), true).
	BorderForeground(color.border).
	Padding(0, 1).
	Render

func (cb ComponentBuilder) NewModuleBox(mod module.Module) string {
	b := newBoxWithTitle()
	width := max(mod.GetWidthOfOutput(), 30)

	return b.Render(mod.Title, mod.Output.String(), width)
}
