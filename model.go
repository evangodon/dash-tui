package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/exp/slices"

	"github.com/evangodon/dash/config"
	"github.com/evangodon/dash/ui"
	"github.com/evangodon/dash/util"
)

type model struct {
	activeTab     int
	activeTabName string
	tabs          []config.Tab
	config        *config.Config
	sub           chan moduleUpdateMsg
	window        ui.Window
	err           error
	openConfig    bool
}

func initialModel(cfg *config.Config, initialTab int) model {

	initialTab = util.Clamp(0, initialTab-1, len(cfg.Tabs)-1)
	return model{
		activeTab:     initialTab,
		activeTabName: cfg.Tabs[initialTab].Name,
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
		if msg.err != nil {
			m.err = msg.err
			return m, tea.Quit
		}

		configErr := m.config.Reload()
		if configErr != nil {
			m.err = configErr
			return m, tea.Quit
		}

		m.openConfig = false
		return m, m.runActiveModules()
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error occured: %v\n", m.err)
	}
	if m.openConfig {
		return ""
	}

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
