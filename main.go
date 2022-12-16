package main

import (
	"flag"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/evangodon/dash/config"
	"github.com/evangodon/dash/dimensions"
)

func main() {
	configPath := flag.String("config", "", "config file")
	initialtab := flag.Int("tab", 1, "index of initial active tab")
	calcdimensions := flag.Bool("dimensions", false, "calculate dimensions of all modules")
	flag.Parse()

	config, err := config.New(*configPath)
	if err != nil {
		logError(err)
	}

	if *calcdimensions {
		p := tea.NewProgram(dimensions.InitialModel(config))
		if _, err := p.Run(); err != nil {
			logError(err)
		}
		os.Exit(0)
	}

	p := tea.NewProgram(initialModel(config, *initialtab), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		logError(err)
	}
}
