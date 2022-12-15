package main

import (
	"log"
	"os"

	"github.com/evangodon/dash/config"
	"github.com/evangodon/dash/ui"
)

func errorTitle(msg string) string {
	return ui.BoldText(ui.RedText(msg))
}

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
