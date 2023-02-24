package main

import (
	"fmt"

	"github.com/evangodon/dash/config"
	"github.com/evangodon/dash/ui"
	"github.com/evangodon/dash/util"
)

func (m model) NewModuleBox(mod config.Module, height int) string {
	b := ui.NewBoxWithTitle()
	content := mod.Output.String()
	title := mod.GetTitle()

	if mod.Err != nil {
		content = ui.RedText(mod.Err.String())
	}
	if mod.Status() == config.StatusLoading {
		title = fmt.Sprintf("%s %s", title, m.spinner.View())
	}
	width := util.Clamp(len(mod.Title), mod.GetRenderedWidth(), m.window.width-4)

	return b.Render(title, content, width, height)
}
