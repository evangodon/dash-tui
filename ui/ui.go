package ui

type Window struct {
	Height int
	Width  int
}

type ComponentBuilder struct {
	window      Window
	titleHeight int
	tabbar      int
}

const (
	titleHeight   = 2
	tabbarHeight  = 3
	helpbarheight = 1
)

func New(w Window) *ComponentBuilder {
	return &ComponentBuilder{
		window:      w,
		titleHeight: titleHeight,
		tabbar:      tabbarHeight,
	}
}
