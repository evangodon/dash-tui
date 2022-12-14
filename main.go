package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/evangodon/dash/config"
)

func main() {

	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	p := tea.NewProgram(initialModel(config), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}

}
