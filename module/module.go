package module

import (
	"bytes"
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

func (m *Module) GetOutputHeight() int {
	if m.Output == nil {
		return 2
	}
	return lg.Height(m.Output.String())
}
