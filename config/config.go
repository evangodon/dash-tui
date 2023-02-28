package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/evangodon/dash/ui"
	"golang.org/x/exp/slices"
)

type Config struct {
	Tabs         []Tab     `toml:"tab"`
	Modules      []*Module `toml:"module"`
	Dependencies []string  `toml:"dependencies"`
	FilePath     string
	Settings     Settings `toml:"settings"`
}

type Settings struct {
	PrimaryColor string `toml:"primary-color"`
}

func New(configPath string) (*Config, error) {

	var cfg = Config{
		Tabs:     []Tab{},
		Modules:  []*Module{},
		FilePath: configPath,
	}

	if err := cfg.ReadConfig(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Config) ReadConfig() error {
	f, err := ioutil.ReadFile(cfg.FilePath)
	if err != nil {
		return err
	}
	cfg.Modules = make([]*Module, 0)
	cfg.Tabs = make([]Tab, 0)

	err = toml.Unmarshal(f, &cfg)
	if err != nil {
		return &ConfigError{Reason: err.Error()}
	}

	if err := cfg.Validate(); err != nil {
		return err
	}

	for _, mod := range cfg.Modules {
		if mod.Dir == "" {
			mod.Dir = filepath.Dir(cfg.FilePath)
		}
		if mod.Dir == "." {
			mod.Dir = cfg.getCWD()
		}
	}

	if cfg.Settings.PrimaryColor == "" {
		cfg.Settings.PrimaryColor = ui.Color.Primary.Dark
	}

	return nil
}

func (cfg *Config) getCWD() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path
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
	// Validate Dependencies
	for _, dep := range cfg.Dependencies {
		_, err := exec.LookPath(dep)
		if err != nil {
			return &ConfigError{
				Reason: err.Error(),
			}
		}
	}

	if len(cfg.Tabs) == 0 {
		return &ConfigError{
			Reason: "need at least one tab in config file",
		}
	}

	// Validate Tabs
	for i, tab := range cfg.Tabs {
		if tab.Name == "" {
			r := fmt.Sprintf("tab at index %d needs a name of length of 1 or more", i)
			return &ConfigError{
				Reason: r,
			}
		}

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
					"Module \"%s\" for tab \"%s\" defined but none exist in config file.\nModules Found: %v",
					modName,
					tab.Name,
					strings.Join(foundModules, ", "),
				)
				return &ConfigError{
					Reason: reason,
				}
			}
		}
	}

	// Validate Settings
	if primaryColor := cfg.Settings.PrimaryColor; primaryColor != "" {
		if len(primaryColor) != 7 || primaryColor[0] != '#' {
			return &ConfigError{
				Reason: fmt.Sprintf("Colors should be of format #ffffff, got %s", primaryColor),
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
