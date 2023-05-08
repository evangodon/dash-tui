package config

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"os"
	"os/exec"
	"time"

	lg "github.com/charmbracelet/lipgloss"
	"github.com/creack/pty"

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
	cmd.Dir = m.Dir
	m.status = StatusLoading

	tty, err := pty.Start(cmd)
	if err != nil {
		panic(err)
	}
	defer tty.Close()

	go func() {
		buf := bufio.NewReader(tty)

		for {
			line, _, err := buf.ReadLine()

			if err == io.EOF {
				break
			}

			if err != nil {
				// Works on Mac but input/output error on Linux
				// log.Fatal(err.Error())
				break
			}

			line = cleanOutput(line)
			m.Output.Write(line)
			m.Output.Write([]byte("\n"))
		}
	}()
	cmd.Wait()

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

func cleanOutput(b []byte) []byte {
	new := bytes.TrimSpace(b)
	new = bytes.Replace(new, []byte("[?1h"), []byte(""), -1)
	return new
}
