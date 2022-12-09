package main

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"

	"github.com/evangodon/dash/module"
)

// todo, split out toml config with app config?
type config struct {
	Tabs     []string
	Modules  map[string]*module.Module
	filePath string
}

func newConfig(configPath string) *config {
	var cfg = config{
		Tabs:     []string{},
		Modules:  map[string]*module.Module{},
		filePath: configPath,
	}

	if !cfg.configFileExists() {
		cfg.createConfigFile()
	}

	f, err := ioutil.ReadFile(cfg.filePath)
	if err != nil {
		panic(err)
	}

	err = toml.Unmarshal(f, &cfg)
	if err != nil {
		panic(err)
	}

	if len(cfg.Tabs) == 0 {
		panic("Need at least one tab in config file")
	}

	return &cfg
}

func (cfg *config) configFileExists() bool {
	_, err := os.Stat(cfg.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return true
}

func (cfg *config) createConfigFile() {
	if cfg.configFileExists() {
		panic("config file already exists")
	}

	file, err := os.Create(cfg.filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func (cfg *config) reload() {
	f, err := ioutil.ReadFile(cfg.filePath)
	if err != nil {
		panic(err)
	}

	err = toml.Unmarshal(f, &cfg)
	if err != nil {
		panic(err)
	}
}

func (cfg *config) RunModules() {
	for _, m := range cfg.Modules {
		m.Run()
	}
}
