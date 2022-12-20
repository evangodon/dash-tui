package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"

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
	configOpen    bool
	help          help.Model
}

func initialModel(cfg *config.Config, initialTab int) model {

	initialTab = util.Clamp(0, initialTab-1, len(cfg.Tabs)-1)
	return model{
		activeTab:     initialTab,
		activeTabName: cfg.Tabs[initialTab].Name,
		tabs:          cfg.Tabs,
		config:        cfg,
		sub:           make(chan moduleUpdateMsg),
		help:          help.New(),
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.waitForModuleUpdate(),
		m.runActiveModules(runOptions{}),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case moduleUpdateMsg:
		return m, m.waitForModuleUpdate()

	case tea.WindowSizeMsg:
		m.window.Height = msg.Height
		m.window.Width = msg.Width
		m.help.Width = msg.Width

	case tea.KeyMsg:
		return m.handleKey(msg)

	case configUpdateMsg:
		if msg.err != nil {
			m.err = msg.err
			return m, tea.Quit
		}

		configErr := m.config.ReadConfig()
		if configErr != nil {
			m.err = configErr
			return m, tea.Quit
		}

		m.configOpen = false
		return m, m.runActiveModules(runOptions{})
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error occured: %v\n", m.err)
	}
	if m.configOpen {
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

	helpBar := m.help.View(keys)
	doc.WriteString("\n")
	doc.WriteString(helpBar)

	return doc.String()
}
