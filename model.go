package main

import (
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"

	"github.com/evangodon/dash/module"
	"github.com/evangodon/dash/ui"
)

type model struct {
	activeTab     int
	activeTabName string
	tabs          []string
	config        *config
	sub           chan module.Module
	window        ui.Window
}

type modch chan module.Module

func initialModel() model {
	cfg := newConfig()
	return model{
		activeTab:     0,
		activeTabName: cfg.Tabs[0],
		tabs:          cfg.Tabs,
		config:        cfg,
		sub:           make(modch),
	}
}

// Get modules that will be visible on this tab
func (m model) getActiveModules() []*module.Module {
	allModules := m.config.Modules
	activeModules := []*module.Module{}

	keys := make([]string, len(m.config.Modules))
	i := 0
	for k := range m.config.Modules {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	for _, k := range keys {
		if allModules[k].Tab == m.activeTabName {
			activeModules = append(activeModules, allModules[k])
		}
	}
	return activeModules
}

func runAllModules(m model) tea.Cmd {
	return func() tea.Msg {
		activeModules := m.getActiveModules()

		for _, activeMod := range activeModules {
			if activeMod.Output == nil {
				go func(activeMod *module.Module) {
					activeMod.Run()
					m.sub <- *activeMod
				}(activeMod)
			}
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
	case tea.WindowSizeMsg:
		m.window.Height = msg.Height
		m.window.Width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		// TODO: clean up tab cases
		case "tab":
			if m.activeTab < len(m.tabs)-1 {
				m.activeTab++
			} else {
				m.activeTab = 0
			}
			m.activeTabName = m.tabs[m.activeTab]
			return m, runAllModules(m)
		case "shift+tab":
			if m.activeTab > 0 {
				m.activeTab--
			} else {
				m.activeTab = len(m.tabs) - 1
			}
			m.activeTabName = m.tabs[m.activeTab]
			return m, runAllModules(m)
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	cb := ui.New(m.window)
	doc := strings.Builder{}
	doc.WriteString(ui.AppTitle("Dash tui"))
	doc.WriteString("\n\n")
	doc.WriteString(cb.BuildTabs(m.activeTab, m.tabs...))

	activeModules := m.getActiveModules()

	boxes := make([]string, len(activeModules))
	for _, mod := range activeModules {
		if mod != nil {
			boxes = append(boxes, cb.NewModuleBox(*mod))
		}
	}

	doc.WriteString(lg.JoinHorizontal(lg.Top, boxes...))

	return doc.String()
}
