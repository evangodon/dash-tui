package config

import (
	"fmt"
	"strings"
)

type ConfigError struct {
	reason string
}

func (ConfigError) Title() string {
	return "Config Error"
}

func (cfgErr *ConfigError) Error() string {
	return cfgErr.reason
}

type ModuleError struct {
	Exitcode int
	output   string
}

func (modErr *ModuleError) Error() string {
	return modErr.output
}

func (modErr *ModuleError) String() string {
	s := strings.Builder{}
	s.WriteString("Module Error\n")
	if modErr.Exitcode != 0 {
		s.WriteString(fmt.Sprintf("Exit Code: %d\n", modErr.Exitcode))
	}
	s.WriteString(modErr.output)

	return s.String()
}
