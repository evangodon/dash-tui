package dimensions

import (
	"strconv"
	"sync"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/evangodon/dash/config"
)

type moduleUpdateMsg struct{}
type modulesDone struct{}

type model struct {
	config  *config.Config
	sub     chan moduleUpdateMsg
	spinner spinner.Model
}

func InitialModel(cfg *config.Config) model {
	s := spinner.New()
	s.Spinner = spinner.Line

	return model{
		config:  cfg,
		sub:     make(chan moduleUpdateMsg),
		spinner: s,
	}
}

func (m model) runAllModules() tea.Msg {
	var wg sync.WaitGroup
	for _, mod := range m.config.Modules {
		wg.Add(1)
		go func(mod *config.Module) {
			mod.Run()
			m.sub <- moduleUpdateMsg{}
			wg.Done()
		}(mod)
	}

	wg.Wait()
	return modulesDone{}
}

func (m model) waitForModuleUpdate() tea.Msg {
	return <-m.sub
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.waitForModuleUpdate,
		m.runAllModules,
		m.spinner.Tick,
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case moduleUpdateMsg:
		return m, m.waitForModuleUpdate
	case modulesDone:
		return m, tea.Quit
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

var container = lipgloss.NewStyle().Padding(0, 2).Render

func (m model) View() string {
	moduleCount := len(m.config.Modules)
	columns := []table.Column{
		{Title: "Status", Width: 8},
		{Title: "Module Name", Width: 30},
		{Title: "Width", Width: 8},
		{Title: "Height", Width: 8},
	}

	rows := make([]table.Row, moduleCount)

	for i, mod := range m.config.Modules {
		icon := m.spinner.View()
		if mod.Output.String() != "" {
			icon = "✓"
		}
		name := mod.Name
		width := strconv.Itoa(mod.GetRenderedWidth())
		height := strconv.Itoa(mod.GetOutputHeight())

		rows[i] = table.Row{icon, name, width, height}
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithHeight(moduleCount+1),
	)

	s := table.DefaultStyles()
	s.Selected = lipgloss.NewStyle()
	t.SetStyles(s)

	return container(t.View())
}
