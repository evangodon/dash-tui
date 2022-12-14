package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/evangodon/dash/config"
)

func main() {
	config, err := config.New()
	if err != nil {
		logError(err)
	}

	p := tea.NewProgram(initialModel(config), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		logError(err)
	}
}
