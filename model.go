package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/evangodon/dash/config"
	"github.com/evangodon/dash/util"
)

type window struct {
	width  int
	height int
}

type model struct {
	activeTab     int
	activeTabName string
	tabs          []config.Tab
	config        *config.Config
	modulesCh     chan moduleUpdateMsg
	tabChangeCh   chan tabChangeMsg
	window        window
	err           error
	configOpen    bool
	help          help.Model
	spinner       spinner.Model
}

func initialModel(cfg *config.Config, initialTab int) model {
	s := spinner.New()
	initialTab = util.Clamp(0, initialTab-1, len(cfg.Tabs)-1)

	return model{
		activeTab:     initialTab,
		activeTabName: cfg.Tabs[initialTab].Name,
		tabs:          cfg.Tabs,
		config:        cfg,
		modulesCh:     make(chan moduleUpdateMsg),
		tabChangeCh:   make(chan tabChangeMsg),
		help:          help.New(),
		spinner:       s,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.waitForModuleUpdate,
		m.runActiveModules(runOptions{}),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case moduleUpdateMsg:
		return m, m.waitForModuleUpdate
	case modulesDoneMsg:
		m.spinner = spinner.New()
		return m, nil

	case tea.WindowSizeMsg:
		m.window.height = msg.Height
		m.window.width = msg.Width
		m.help.Width = msg.Width
		return m, nil

	case tea.KeyMsg:
		return m.handleKey(msg)

	case configUpdateMsg:
		if msg.err != nil {
			m.err = msg.err
			return m, nil
		}

		configErr := m.config.ReadConfig()
		if configErr != nil {
			m.err = configErr
			return m, nil
		}

		m.configOpen = false
		return m, m.runActiveModules(runOptions{})

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	if m.err != nil {
		return styleError(m.err) + "\n Press e to edit config\n Press q to quit"
	}
	if m.configOpen {
		return ""
	}

	doc := strings.Builder{}
	doc.WriteString(m.BuildTabs(m.activeTab, m.tabs...))
	doc.WriteString("\n")

	activeModules := m.activeModules()
	tab := m.NewGrid(activeModules)
	doc.WriteString(tab)

	helpBar := m.help.View(keys)
	doc.WriteString("\n")
	doc.WriteString(helpBar)

	return doc.String()
}
