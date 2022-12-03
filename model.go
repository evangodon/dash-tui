package main

import (
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/evangodon/dash/module"
	"github.com/evangodon/dash/ui"
)

type model struct {
	activeTab int
	tabs      []string
	config    *config
	sub       chan module.Module
}

type modch chan module.Module

func initialModel() model {
	cfg := newConfig()
	return model{
		activeTab: 0,
		tabs:      cfg.TabsList,
		config:    cfg,
		sub:       make(modch),
	}
}

func runAllModules(m model) tea.Cmd {
	return func() tea.Msg {
		for _, termMod := range m.config.Modules {
			go func(termMod *module.Module) {
				termMod.Run()
				m.sub <- *termMod
			}(termMod)
		}
		return nil
	}
}

func waitForModuleUpdate(sub modch) tea.Cmd {
	return func() tea.Msg {
		return <-sub
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		waitForModuleUpdate(m.sub),
		runAllModules(m),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case module.Module:
		return m, waitForModuleUpdate(m.sub)
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

	keys := make([]string, len(m.config.Modules))
	i := 0
	for k := range m.config.Modules {
		keys = append(keys, k)
		i++
	}

	sort.Strings(keys)

	for _, k := range keys {
		module := m.config.Modules[k]

		if module != nil {
			doc.WriteString(module.Output.String())
		}
	}

	return doc.String()
}
