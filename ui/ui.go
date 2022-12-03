package ui

type Window struct {
	Height int
	Width  int
}

type ComponentBuilder struct {
	window Window
}

func New(w Window) *ComponentBuilder {
	return &ComponentBuilder{
		window: w,
	}
}
