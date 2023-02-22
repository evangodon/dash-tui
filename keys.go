package main

import (
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type keyMap struct {
	NextTab       key.Binding
	PrevTab       key.Binding
	EditConfig    key.Binding
	ReloadConfig  key.Binding
	ReloadModules key.Binding
	Help          key.Binding
	Quit          key.Binding
}

var keys = keyMap{
	NextTab: key.NewBinding(
		key.WithKeys("tab", "l"),
		key.WithHelp("l", "next tab"),
	),
	PrevTab: key.NewBinding(
		key.WithKeys("shift+tab", "h"),
		key.WithHelp("h", "prev tab"),
	),
	ReloadModules: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "reload modules"),
	),
	EditConfig: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "edit config"),
	),
	ReloadConfig: key.NewBinding(
		key.WithKeys("R"),
		key.WithHelp("R", "reload config"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.PrevTab, k.NextTab, k.EditConfig, k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.PrevTab, k.NextTab},
		{k.ReloadModules, k.ReloadConfig},
		{k.EditConfig, k.Help},
		{k.Quit},
	}
}

func (m model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, keys.EditConfig):
		m.configOpen = true
		m.err = nil
		return m, m.openConfigInEditor()
	case key.Matches(msg, keys.ReloadModules):
		return m, m.runActiveModules(runOptions{force: true})
	case key.Matches(msg, keys.ReloadConfig):
		configErr := m.config.ReadConfig()
		if configErr != nil {
			m.err = configErr
			return m, tea.Quit
		}
		return m, m.runActiveModules(runOptions{force: true})
	case key.Matches(msg, keys.NextTab):
		m.activeTab = (m.activeTab + 1) % len(m.tabs)
		m.activeTabName = m.tabs[m.activeTab].Name
		return m, m.runActiveModules(runOptions{})
	case key.Matches(msg, keys.PrevTab):
		m.activeTab--
		if m.activeTab < 0 {
			m.activeTab = len(m.tabs) - 1
		}
		m.activeTabName = m.tabs[m.activeTab].Name
		return m, m.runActiveModules(runOptions{})
	case key.Matches(msg, keys.Help):
		m.help.ShowAll = !m.help.ShowAll
		return m, nil
	case key.Matches(msg, keys.Quit):
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
