package config

type service struct {
	CheckInterval int `json:"check-interval" mapstructure:"check-interval"`
}

type manager struct {
	Address string `json:"address" mapstructure:"address"`
}

type state struct {
	Address string `json:"address" mapstructure:"address"`
}

type Config struct {
	Manager manager `json:"manager" mapstructure:"manager"`
	State   state   `json:"state" mapstructure:"state"`
	Service service `json:"service" mapstructure:"service"`
}

// Load initializes a default configuration struct.
func Load() *Config {
	return &Config{}
}
