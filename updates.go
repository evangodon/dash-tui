package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/evangodon/dash/config"
	"golang.org/x/exp/slices"
)

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

type runOptions struct {
	force bool
}

func (m model) runActiveModules(options runOptions) tea.Cmd {
	return func() tea.Msg {
		activeModules := m.getActiveModules()

		for _, activeMod := range activeModules {
			if activeMod.Output == nil || options.force {
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
