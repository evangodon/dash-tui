package ui

import (
	"github.com/evangodon/dash/config"
)

func (cb ComponentBuilder) NewModuleBox(mod config.Module, height int) string {
	b := newBoxWithTitle()

	return b.Render(mod.GetTitle(), mod.Output.String(), mod.GetRenderedWidth(), height)
}
