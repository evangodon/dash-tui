package config

type ConfigError struct {
	reason string
}

func (cfgErr *ConfigError) Error() string {
	return cfgErr.reason
}
