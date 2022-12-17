package util

import (
	"os"
	"os/exec"
)

func OpenFileInEditor(path string) error {

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}

	e := exec.Command(editor, path)
	e.Stdin = os.Stdin
	e.Stdout = os.Stdout
	e.Stderr = os.Stderr

	return e.Run()
}
