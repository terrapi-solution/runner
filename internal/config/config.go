package config

// Config defines the general configuration
var globalConfig *Config

// Config defines the service configuration.
type service struct {
	CheckInterval int `mapstructure:"checkInterval"`
}

// controller defines the controller configuration.
type controller struct {
	Host         string       `mapstructure:"host"`
	Port         int          `mapstructure:"port"`
	Certificates certificates `mapstructure:"certificates"`
}

// certificates defines the controller credential configuration.
type certificates struct {
	Status   bool   `mapstructure:"status"`
	CertFile string `mapstructure:"cert"`
	KeyFile  string `mapstructure:"key"`
	CaFile   string `mapstructure:"ca"`
}

// state defines the state configuration.
type state struct {
	Host  string `mapstructure:"host"`
	Port  int    `mapstructure:"port"`
	Token string `mapstructure:"token"`
}

// Config defines the general configuration.
type Config struct {
	Controller controller `mapstructure:"controller"`
	State      state      `mapstructure:"state"`
	Service    service    `mapstructure:"service"`
}

// Load initializes a default configuration struct.
func Load() *Config {
	return &Config{}
}

// Get returns the global configuration.
func Get() *Config { return globalConfig }

// Set returns the global configuration.
func Set(config *Config) { globalConfig = config }
