package ui

import (
	"strings"

	lg "github.com/charmbracelet/lipgloss"

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

func (cb ComponentBuilder) NewTabLayout(modules []*module.Module) string {

	rows := []row{}
	currentRow := row{}
	borderWidth := 2
	for _, mod := range modules {
		boxwidth := Max(mod.GetRenderedWidth(), len(mod.Title))
		boxheight := mod.GetOutputHeight()

		if cb.window.Width > currentRow.width+boxwidth+len(currentRow.items) {
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
			gap := " "
			newbox := cb.NewModuleBox(*item, row.height)
			boxes = append(boxes, gap, newbox)
		}

		rendered := lg.JoinHorizontal(lg.Top, boxes...)
		doc.WriteString(rendered)
		doc.WriteString("\n")
	}

	height := cb.window.Height - titleHeight - tabbarHeight
	s := doc.String()
	container := lg.NewStyle().MaxHeight(height)

	return container.Render(s)
}
