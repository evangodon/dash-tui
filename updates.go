package main

import (
	"context"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/evangodon/dash/config"
	"golang.org/x/exp/slices"
)

// Get modules that will be visible on this tab
func (m model) activeModules() []*config.Module {
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

// Start spinner ticking and run all modules
func (m model) runActiveModules(options runOptions) tea.Cmd {
	var wg sync.WaitGroup

	return tea.Batch(m.spinner.Tick, func() tea.Msg {
		activeModules := m.activeModules()
		for _, activeMod := range activeModules {
			if activeMod.Output == nil || options.force {
				wg.Add(1)
				ctx, cancel := context.WithTimeout(context.Background(), config.CmdTimeout)
				defer cancel()
				done := make(chan bool, 1)

				go func(activeMod *config.Module) {
					activeMod.Run(ctx, done)
					for {
						select {
						case <-done:
							wg.Done()
							m.modulesCh <- moduleUpdateMsg{}
							break
						// Cancel module execution if tab changes
						case <-m.tabChangeCh:
							cancel()
							break
						}
					}
				}(activeMod)
			}
		}
		wg.Wait()
		return modulesDoneMsg{}
	})
}

func (m model) waitForModuleUpdate() tea.Msg {
	return <-m.modulesCh
}
