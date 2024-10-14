package cmd

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var (
	managerServerAddr = "localhost:8080"
	stateServerAddr   = "localhost:8080"
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

// Returns a context with a signal handler for graceful shutdown
func getContext() context.Context {
	ctx := context.Background()

	// Trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

	return ctx
}
