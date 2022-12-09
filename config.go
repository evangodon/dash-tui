package main

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"

	"github.com/evangodon/dash/module"
)

type config struct {
	Tabs     []string
	Modules  map[string]*module.Module
	filePath string
}

func newConfig() *config {
	var cfg = config{
		Tabs:     []string{},
		Modules:  map[string]*module.Module{},
		filePath: "./config.toml",
	}
	f, err := ioutil.ReadFile(cfg.filePath)
	if err != nil {
		panic(err)
	}

	err = toml.Unmarshal(f, &cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}

func (cfg *config) RunModules() {
	for _, m := range cfg.Modules {
		m.Run()
	}
}
