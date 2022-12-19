package ui

import (
	"github.com/evangodon/dash/config"
)

func (ComponentBuilder) NewModuleBox(mod config.Module, height int) string {
	b := newBoxWithTitle()
	content := mod.Output.String()
	if mod.Err != nil {
		content = mod.Err.String()
	}

	return b.Render(mod.GetTitle(), content, mod.GetRenderedWidth(), height)
}
