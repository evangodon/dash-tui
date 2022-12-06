package main

import tea "github.com/charmbracelet/bubbletea"

func (m model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
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
	default:
		return m, nil
	}
}
