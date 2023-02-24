package config

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"time"

	lg "github.com/charmbracelet/lipgloss"

	"github.com/evangodon/dash/util"
)

type Module struct {
	Name   string
	Title  string
	Exec   string
	Output *bytes.Buffer
	Width  int
	Err    *ModuleError
	Dir    string
	status Status
}

type Status int

const (
	StatusInitial Status = iota
	StatusLoading
	StatusFinished
)

const cmdTimeout = 8 * time.Second

func (m *Module) GetTitle() string {
	if m.Title != "" {
		return m.Title
	}
	return m.Name
}

func (m *Module) Run() {
	m.Output = new(bytes.Buffer)

	ctx, cancel := context.WithTimeout(context.Background(), cmdTimeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, "bash", "-c", m.Exec)

	cmd.Env = os.Environ()
	// Preserve ANSI Colors
	cmd.Env = append(cmd.Env, "CLICOLOR_FORCE=1", "GH_FORCE_TTY=true")
	cmd.Stdout = m.Output
	cmd.Stderr = m.Output
	cmd.Dir = m.Dir

	m.status = StatusLoading

	err := cmd.Run()

	m.status = StatusFinished
	if err != nil {
		var code int
		if exiterr, ok := err.(*exec.ExitError); ok {
			code = exiterr.ExitCode()
		}

		output := m.Output.String()
		if output == "" {
			output = err.Error()
		}

		m.Err = &ModuleError{
			Exitcode: code,
			output:   output,
		}
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
	if m.Width > 0 {
		return m.Width
	}
	titleLength := len(m.GetTitle())
	moduleWidth := lg.Width(m.Output.String())

	return util.Max(titleLength, moduleWidth) + borderWidth + paddingWidth
}

func (m *Module) GetOutputHeight() int {
	if m.Output == nil {
		return 2
	}
	return lg.Height(m.Output.String())
}

func (m *Module) Status() Status {
	return m.status
}
