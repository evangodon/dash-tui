package config

import (
	"bytes"
	"os"
	"os/exec"

	lg "github.com/charmbracelet/lipgloss"
)

type Module struct {
	Name   string
	Title  string
	Exec   string
	Output *bytes.Buffer
	Width  int
	error  error
}

func (m *Module) GetTitle() string {
	if m.Title != "" {
		return m.Title
	}
	return m.Name
}

func (m *Module) Run() {
	m.Output = new(bytes.Buffer)
	cmd := exec.Command("sh", "-c", m.Exec)

	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "CLICOLOR_FORCE=1")
	cmd.Stdout = m.Output
	cmd.Stderr = m.Output

	err := cmd.Run()
	if err != nil {
		m.error = err
	}
}

// GetWidthOfOutput returns the width of the output after running, not the actual box
func (m *Module) GetWidthOfOutput() int {
	if m.Output == nil {
		return 0
	}
	return lg.Width(m.Output.String())
}

var (
	borderWidth  = 2
	paddingWidth = 2
)

// GetRenderedWidth returns the actual width that the module will take
func (m *Module) GetRenderedWidth() int {
	if m.Output == nil {
		return len(m.GetTitle()) + borderWidth + paddingWidth
	}
	if m.Width > 0 {
		return m.Width
	}
	return lg.Width(m.Output.String()) + borderWidth + paddingWidth
}

func (m *Module) GetOutputHeight() int {
	if m.Output == nil {
		return 2
	}
	return lg.Height(m.Output.String())
}