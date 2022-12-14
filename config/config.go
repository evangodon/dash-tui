package config

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/adrg/xdg"
)

type Config struct {
	Tabs     []Tab     `toml:"tab"`
	Modules  []*Module `toml:"module"`
	FilePath string
}

var (
	appName = "dashtui"
)

func New() (*Config, error) {
	defaultConfigPath, err := xdg.ConfigFile(appName + "/config.toml")
	if err != nil {
		log.Fatal(err)
	}
	configPath := flag.String("config", defaultConfigPath, "config file")
	flag.Parse()

	var cfg = Config{
		Tabs:     []Tab{},
		Modules:  []*Module{},
		FilePath: *configPath,
	}

	if *configPath != defaultConfigPath && !cfg.configFileExists() {
		return nil, errors.New("config file not found")
	}

	if !cfg.configFileExists() {
		cfg.mustCreateConfigFile()
	}

	f, err := ioutil.ReadFile(cfg.FilePath)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(f, &cfg)
	if err != nil {
		return nil, err
	}

	if len(cfg.Tabs) == 0 {
		panic("Need at least one tab in config file")
	}

	return &cfg, nil
}

func (cfg *Config) configFileExists() bool {
	_, err := os.Stat(cfg.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return true
}

func (cfg *Config) mustCreateConfigFile() {
	if cfg.configFileExists() {
		panic("config file already exists")
	}

	file, err := os.Create(cfg.FilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func (cfg *Config) Reload() {
	f, err := ioutil.ReadFile(cfg.FilePath)
	if err != nil {
		panic(err)
	}

	err = toml.Unmarshal(f, &cfg)
	if err != nil {
		panic(err)
	}
	for _, mod := range cfg.Modules {
		mod.Output = nil
	}
}

func (cfg *Config) RunModules() {
	for _, m := range cfg.Modules {
		m.Run()
	}
}
