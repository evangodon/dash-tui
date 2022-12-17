package main

import (
	"flag"
	"os"

	"github.com/adrg/xdg"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/evangodon/dash/config"
	"github.com/evangodon/dash/dimensions"
	"github.com/evangodon/dash/util"
)

var (
	appName = "dashtui"
)

func main() {
	configPath := flag.String("config", "", "config file")
	openConfig := flag.Bool("open-config", false, "open config file")
	initialtab := flag.Int("tab", 1, "index of initial active tab")
	calcdimensions := flag.Bool("dimensions", false, "calculate dimensions of all modules")
	flag.Parse()

	defaultConfigPath, err := xdg.ConfigFile(appName + "/config.toml")
	if err != nil {
		logError(err)
	}
	if *configPath == "" {
		*configPath = defaultConfigPath
	}

	config, err := config.New(*configPath)
	if err != nil {
		logError(err)
	}

	if *openConfig {
		if err := util.OpenFileInEditor(*configPath); err != nil {
			logError(err)
		}
		os.Exit(0)
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
