package command

import (
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Loads the application configuration using Viper.
func setupConfig() {
	if viper.GetString("config.file") != "" {
		viper.SetConfigFile(viper.GetString("config.file"))
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("/etc/terrapi/runner")
		viper.AddConfigPath("$HOME/.terrapi/")
	}

	viper.SetEnvPrefix("terrapi")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := readConfig(); err != nil {
		log.Error().
			Err(err).
			Msg("Failed to read config file")
	}

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
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		return nil
	}

	// Return nil if there was a file path error
	if _, ok := err.(*os.PathError); ok {
		return nil
	}

	// Return the error for other issues
	return err
}
