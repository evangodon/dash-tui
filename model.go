package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/exp/slices"

	"github.com/evangodon/dash/config"
	"github.com/evangodon/dash/ui"
)

type model struct {
	activeTab     int
	activeTabName string
	tabs          []config.Tab
	config        *config.Config
	sub           chan moduleUpdateMsg
	window        ui.Window
}

func initialModel(cfg *config.Config) model {
	return model{
		activeTab:     0,
		activeTabName: cfg.Tabs[0].Name,
		tabs:          cfg.Tabs,
		config:        cfg,
		sub:           make(chan moduleUpdateMsg),
	}
}

// Get modules that will be visible on this tab
func (m model) getActiveModules() []*config.Module {
	allModules := m.config.Modules

	activeTab := m.tabs[m.activeTab]
	activeModules := make([]*config.Module, len(activeTab.Modules))

	for i, moduleName := range activeTab.Modules {
		idxModule := slices.IndexFunc(allModules, func(mod *config.Module) bool {
			return mod.Name == moduleName
		})

		activeModules[i] = allModules[idxModule]
	}
	return activeModules
}

func (m model) runActiveModules() tea.Cmd {
	return func() tea.Msg {
		activeModules := m.getActiveModules()

		for _, activeMod := range activeModules {
			if activeMod.Output == nil {
				go func(activeMod *config.Module) {
					activeMod.Run()
					m.sub <- moduleUpdateMsg{}
				}(activeMod)
			}
		}
		return nil
	}
}

func (m model) waitForModuleUpdate() tea.Cmd {
	return func() tea.Msg {
		return <-m.sub
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.waitForModuleUpdate(),
		m.runActiveModules(),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case moduleUpdateMsg:
		return m, m.waitForModuleUpdate()

	case tea.WindowSizeMsg:
		m.window.Height = msg.Height
		m.window.Width = msg.Width

	case tea.KeyMsg:
		return m.handleKey(msg)

	case configUpdateMsg:
		m.config.Reload()
		return m, m.runActiveModules()
	}
	return m, nil
}

func (m model) View() string {
	cb := ui.New(m.window)
	doc := strings.Builder{}
	doc.WriteString(ui.AppTitle("Dash TUI"))
	doc.WriteString("\n")
	doc.WriteString(cb.BuildTabs(m.activeTab, m.tabs...))
	doc.WriteString("\n")

	activeModules := m.getActiveModules()

	tab := cb.NewTabLayout(activeModules)
	doc.WriteString(tab)

	return doc.String()
}
