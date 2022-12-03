package module

import (
	"bytes"
	"log"
	"os/exec"
)

type Module struct {
	Title  string
	Tab    string
	File   string
	Output *bytes.Buffer
	Error  error
}

func (m *Module) Run() {
	if m.File == "" {
		log.Fatal("file not set for module: ", m.Title)
	}
	m.Output = new(bytes.Buffer)
	cmd := exec.Command(m.File)

	cmd.Stdout = m.Output
	cmd.Stderr = m.Output

	err := cmd.Run()
	if err != nil {
		log.Fatal("command failed", err)
	}
}
