package command

import (
	"errors"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Loads the application configuration using Viper.
func setupConfig() {
	// Set the default configuration values
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/terrapi/runner")
	viper.AddConfigPath("$HOME/.terrapi/runner")

	// Set the environment variables
	viper.SetEnvPrefix("TERRAPI")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Attempt to read the configuration file
	if err := readConfig(); err != nil {
		log.Error().
			Err(err).
			Msg("Failed to read config file")
	}

	// Unmarshal the configuration into the struct
	if err := viper.Unmarshal(cfg); err != nil {
		log.Error().
			Err(err).
			Msg("Failed to parse config file")
	}
}

// Attempts to read the configuration from the file specified.
func readConfig() error {
	err := viper.ReadInConfig()

	// Return nil if the config was read successfully
	if err == nil {
		return nil
	}

	// Return nil if the config file was not found
	var configFileNotFoundError viper.ConfigFileNotFoundError
	if errors.As(err, &configFileNotFoundError) {
		return nil
	}

	// Return nil if there was a file path error
	var pathError *os.PathError
	if errors.As(err, &pathError) {
		return nil
	}

	// Return the error for other issues
	return err
}
