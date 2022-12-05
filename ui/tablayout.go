package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/evangodon/dash/module"
)

type row struct {
	items  []*module.Module
	width  int
	height int
}

func (r *row) AddModule(module *module.Module) {
	r.items = append(r.items, module)
}

// TODO: clean this up
// TODO: update the module.GetWidth method
func (cb ComponentBuilder) NewTabLayout(modules []*module.Module) string {

	rows := []row{}
	currentRow := row{}
	borderWidth := 2
	for _, mod := range modules {
		boxwidth := max(lipgloss.Width(mod.Output.String()), len(mod.Title))
		boxheight := lipgloss.Height(mod.Output.String())

		if cb.window.Width > currentRow.width+boxwidth+borderWidth {
			currentRow.AddModule(mod)
			currentRow.width += boxwidth + borderWidth
			if boxheight > currentRow.height {
				currentRow.height = boxheight
			}
		} else {
			rows = append(rows, currentRow)
			currentRow = row{
				items:  []*module.Module{mod},
				width:  boxwidth,
				height: boxheight,
			}
		}
	}
	rows = append(rows, currentRow)

	doc := strings.Builder{}
	for _, row := range rows {
		boxes := []string{}
		for _, item := range row.items {
			newbox := cb.NewModuleBox(*item, row.height)
			boxes = append(boxes, newbox)
		}

		rendered := lipgloss.JoinHorizontal(lipgloss.Top, boxes...)
		doc.WriteString(rendered)
		doc.WriteString("\n")
	}

	return doc.String()
}
