package main

import (
	"fmt"
	"strings"

	lg "github.com/charmbracelet/lipgloss"
	"github.com/evangodon/dash/config"
	"github.com/evangodon/dash/ui"
	"github.com/evangodon/dash/util"
)

func (m model) NewModuleBox(mod config.Module, height int) string {
	b := NewBoxWithTitle()
	content := mod.Output.String()
	if mod.Err != nil {
		content = mod.Err.String()
	}
	title := mod.GetTitle()
	if mod.Status() == config.StatusLoading {
		title = fmt.Sprintf("%s %s", title, m.spinner.View())
	}

	return b.Render(title, content, mod.GetRenderedWidth(), height)
}

type BoxWithLabel struct {
	BoxStyle   lg.Style
	LabelStyle lg.Style
}

func NewBoxWithTitle() BoxWithLabel {
	return BoxWithLabel{
		BoxStyle: lg.NewStyle().
			Border(lg.RoundedBorder()).
			BorderForeground(ui.Color.Border).
			Padding(0, 1),

		LabelStyle: lg.NewStyle().
			Foreground(ui.Color.Border).
			PaddingTop(0).
			PaddingBottom(0).
			PaddingLeft(1).
			PaddingRight(1),
	}
}

func (b BoxWithLabel) Render(label, content string, width int, height int) string {
	var (
		border          = b.BoxStyle.GetBorderStyle()
		topBorderStyler = lg.NewStyle().
				Foreground(b.BoxStyle.GetBorderTopForeground()).
				Render
		topLeft  = topBorderStyler(border.TopLeft)
		topRight = topBorderStyler(border.TopRight)

		renderedLabel = b.LabelStyle.Render(label)
	)

	borderWidth := b.BoxStyle.GetHorizontalBorderSize()
	cellsShort := util.Max(0, width+borderWidth-lg.Width(topLeft+topRight+renderedLabel))
	gap := strings.Repeat(border.Top, cellsShort)
	top := topLeft + renderedLabel + topBorderStyler(gap) + topRight

	bottom := b.BoxStyle.Copy().
		BorderTop(false).
		Width(width).
		Height(height).
		Render(content)

	return top + "\n" + bottom
}
