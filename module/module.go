package module

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
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
		log.Fatal("file not set for module: ", m.Title)
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

	lines := strings.Split(m.Output.String(), "\n")

	longest := lines[0]
	for _, s := range lines {
		if len(s) > len(longest) {
			longest = s
		}
	}

	return len(longest)
}
