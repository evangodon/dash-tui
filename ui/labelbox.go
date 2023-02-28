package ui

import (
	"strings"

	lg "github.com/charmbracelet/lipgloss"
	"github.com/evangodon/dash/util"
)

type BoxWithLabel struct {
	BoxStyle   lg.Style
	LabelStyle lg.Style
}

type Option func(*BoxWithLabel)

func WithBoxStyle(boxstyle lg.Style) Option {
	return func(box *BoxWithLabel) {
		box.BoxStyle = boxstyle
	}
}

func NewBoxWithTitle(opts ...Option) BoxWithLabel {
	box := BoxWithLabel{
		BoxStyle: lg.NewStyle().
			Border(lg.RoundedBorder()).
			BorderForeground(Color.Primary).
			Padding(0, 1),

		LabelStyle: lg.NewStyle().
			Foreground(Color.Border).
			PaddingTop(0).
			PaddingBottom(0).
			PaddingLeft(1).
			PaddingRight(1),
	}
	for _, opt := range opts {
		opt(&box)
	}
	return box
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
