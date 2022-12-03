package main

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"

	"github.com/evangodon/dash/module"
)

type config struct {
	TabsList []string
	Modules  map[string]*module.Module
}

func newConfig() *config {
	var cfg config
	f, err := ioutil.ReadFile("./config.toml")
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
