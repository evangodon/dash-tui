package main

import (
	"flag"
	"fmt"
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

var defaultConfigPath string
var configErr error

func init() {
	defaultConfigPath, configErr = xdg.ConfigFile(appName + "/config.toml")
	if configErr != nil {
		logError(configErr)
	}
}

func main() {
	configPath := flag.String("config", defaultConfigPath, "config file")
	openConfig := flag.Bool("open-config", false, "open config file")
	initialtab := flag.Int("tab", 1, "index of initial active tab")
	calcdimensions := flag.Bool("dimensions", false, "calculate dimensions of all modules")
	flag.Parse()

	if exists := util.FileExists(*configPath); !exists {
		r := fmt.Sprintf("config file not found at \"%s\"", *configPath)
		logError(&config.ConfigError{Reason: r})
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
