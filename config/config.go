package config

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/adrg/xdg"
	"golang.org/x/exp/slices"
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
		r := fmt.Sprintf("no config file found at \"%s\"", *configPath)
		return nil, &ConfigError{reason: r}
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
		return nil, &ConfigError{reason: err.Error()}
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
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

func (cfg *Config) Validate() *ConfigError {
	if len(cfg.Tabs) == 0 {
		return &ConfigError{
			reason: "need at least one tab in config file",
		}
	}

	for _, tab := range cfg.Tabs {
		for _, modName := range tab.Modules {
			_, err := cfg.GetModule(modName)
			if err != nil {
				foundModules := func() []string {
					modules := make([]string, len(cfg.Modules))
					for i, mod := range cfg.Modules {
						modules[i] = mod.Name
					}
					return modules
				}()
				reason := fmt.Sprintf(
					"Looked for module \"%s\" for tab \"%s\" but none exist in config file.\nModules Found %v",
					modName,
					tab.Name,
					foundModules,
				)
				return &ConfigError{
					reason: reason,
				}
			}
		}
	}

	return nil
}

func (cfg *Config) GetModule(moduleName string) (*Module, error) {
	idxModule := slices.IndexFunc(cfg.Modules, func(mod *Module) bool {
		return mod.Name == moduleName
	})

	if idxModule == -1 {
		return nil, errors.New("module not found")
	}

	return cfg.Modules[idxModule], nil
}

func (cfg *Config) Reload() error {
	f, err := ioutil.ReadFile(cfg.FilePath)
	if err != nil {
		return err
	}

	err = toml.Unmarshal(f, &cfg)
	if err != nil {
		return err
	}

	if err := cfg.Validate(); err != nil {
		return err
	}

	for _, mod := range cfg.Modules {
		mod.Output = nil
	}
	return nil
}
