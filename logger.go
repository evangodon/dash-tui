package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/evangodon/dash/config"
	"github.com/evangodon/dash/ui"
)

func errorTitle(msg string) string {
	return ui.BoldText(ui.RedText(msg))
}

var errorContainer = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(ui.ColorError).
	Padding(0, 1).
	Render

func logError(err error) {
	l := log.New(os.Stderr, "", 0)
	var errType string
	var msg string

	switch t := err.(type) {
	case *config.ConfigError:
		errType = errorTitle(t.Title())
		msg = fmt.Sprintf("%s\n%s", errType, t.Error())
		msg = errorContainer(msg)
	default:
		errType := errorTitle("Error")
		msg = fmt.Sprintf("%s\n%s", errType, t.Error())
	}

	l.Fatal(msg)
}
