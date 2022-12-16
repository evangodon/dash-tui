package main

import (
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "e":
		m.openConfig = true
		return m, m.openConfigInEditor()
	case "tab", "l":
		m.activeTab = (m.activeTab + 1) % len(m.tabs)
		m.activeTabName = m.tabs[m.activeTab].Name
		return m, m.runActiveModules()
	case "shift+tab", "h":
		m.activeTab--
		if m.activeTab < 0 {
			m.activeTab = len(m.tabs) - 1
		}
		m.activeTabName = m.tabs[m.activeTab].Name
		return m, m.runActiveModules()
	case "ctrl+c", "q":
		return m, tea.Quit
	default:
		return m, nil
	}
}

func (m model) openConfigInEditor() tea.Cmd {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}
	c := exec.Command(editor, m.config.FilePath)
	return tea.ExecProcess(c, func(err error) tea.Msg {
		return configUpdateMsg{err}
	})
}
