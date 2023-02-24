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

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return true
}
