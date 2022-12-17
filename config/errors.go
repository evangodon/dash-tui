package config

type ConfigError struct {
	reason string
}

func (ConfigError) Title() string {
	return "Config Error"
}

func (cfgErr *ConfigError) Error() string {
	return cfgErr.reason
}
