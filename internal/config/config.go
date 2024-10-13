package config

type managerConfig struct {
	Address string `json:"address" mapstructure:"address"`
}

type stateConfig struct {
	Address string `json:"address" mapstructure:"address"`
}

type Config struct {
	Manager managerConfig `json:"manager" mapstructure:"manager"`
	State   stateConfig   `json:"state" mapstructure:"state"`
}

// Load initializes a default configuration struct.
func Load() *Config {
	return &Config{}
}
