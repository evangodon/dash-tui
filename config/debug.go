package config

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func (cfg *Config) DebugLog(msg any) {
	f, err := tea.LogToFile("./tmp/debug.log", "")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := fmt.Sprintf("\n%v", msg)
	logger := log.New(f, "", 2)
	logger.Println(s)
}
