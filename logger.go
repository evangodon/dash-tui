package main

import (
	"log"
	"os"

	"github.com/charmbracelet/lipgloss"

	"github.com/evangodon/dash/config"
)

var (
	errorTitle = lipgloss.NewStyle().Foreground(lipgloss.Color("#f38ba8")).Bold(true).Render
)

func logError(err error) {
	l := log.New(os.Stdout, "", 0)

	switch t := err.(type) {
	case *config.ConfigError:
		errTitle := errorTitle("ConfigError")
		l.Printf("\n%s %s", errTitle, t.Error())
	default:
		errTitle := errorTitle("Error")
		l.Printf("\n%s %s", errTitle, t.Error())
	}
	os.Exit(1)
}
