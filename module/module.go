package module

import (
	"bytes"
	"fmt"
	"os/exec"

	lg "github.com/charmbracelet/lipgloss"
)

type Module struct {
	Title  string
	Tab    string
	Exec   string
	Output *bytes.Buffer
	Error  error
}

func (m *Module) Run() {
	if m.Exec == "" {
		panic(fmt.Errorf("file not set for module: %v", m.Title))
	}
	m.Output = new(bytes.Buffer)
	cmd := exec.Command("sh", "-c", m.Exec)

	cmd.Stdout = m.Output
	cmd.Stderr = m.Output

	err := cmd.Run()
	if err != nil {
		m.Error = err
	}
}

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

func (m *Module) GetRenderedWidth() int {
	if m.Output == nil {
		return 0
	}
	return lg.Width(m.Output.String()) + borderWidth + paddingWidth
}

func (m *Module) GetRenderedHeight() int {
	if m.Output == nil {
		return 2
	}
	return lg.Height(m.Output.String()) + borderWidth + paddingWidth
}
