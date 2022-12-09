package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/adrg/xdg"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	appName = "dashtui"
)

func main() {

	defaultConfig, err := xdg.ConfigFile(appName + "/config.toml")
	if err != nil {
		log.Fatal(err)
	}
	configPath := flag.String("config", defaultConfig, "config file")
	flag.Parse()

	config := newConfig(*configPath)

	p := tea.NewProgram(initialModel(config), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}

}
