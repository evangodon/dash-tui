package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/evangodon/dash/ui"
)

type model struct {
	activeTab int
	tabs      []string
	config    *config
}

func initialModel() model {
	cfg := newConfig()
	for _, m := range cfg.Modules {
		m.Run()
	}

	return model{
		activeTab: 0,
		tabs:      cfg.TabsList,
		config:    cfg,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			if m.activeTab < len(m.tabs)-1 {
				m.activeTab++
			} else {
				m.activeTab = 0
			}

			return m, nil
		case "shift+tab":
			if m.activeTab > 0 {
				m.activeTab--
			} else {
				m.activeTab = len(m.tabs) - 1
			}

			return m, nil
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	components := ui.New()

	doc := strings.Builder{}
	doc.WriteString(components.BuildTabs(m.activeTab, m.tabs...))
	for _, m := range m.config.Modules {
		doc.WriteString(m.Title + ":\n")
		doc.WriteString(m.Output.String())
	}

	return doc.String()
}
