package cmd

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/terrapi-solution/runner/internal/watcher"
	"time"
)

// Command-line flags default values
var (
	serviceCmd = &cobra.Command{
		Use:   "service",
		Short: "Start runner as service",
		Run:   serviceAction,
		Args:  cobra.NoArgs,
	}

	defaultCheckInterval = 5
)

// Initialization of CLI flags and viper config binding
func init() {
	startCmd.AddCommand(serviceCmd)

	//region global parameters
	serviceCmd.PersistentFlags().String("manager-server", managerServerAddr, "Address of the runner manager")
	viper.SetDefault("manager.address", managerServerAddr)
	_ = viper.BindPFlag("manager.address", serviceCmd.PersistentFlags().Lookup("manager-server"))

	serviceCmd.PersistentFlags().String("state-server", stateServerAddr, "Address of the state manager")
	viper.SetDefault("state.address", stateServerAddr)
	_ = viper.BindPFlag("state.address", serviceCmd.PersistentFlags().Lookup("state-server"))
	//endregion

	serviceCmd.PersistentFlags().Int("check-interval", defaultCheckInterval, "Address of the state manager")
	viper.SetDefault("service.check-interval", defaultCheckInterval)
	_ = viper.BindPFlag("service.check-interval", serviceCmd.PersistentFlags().Lookup("check-interval"))
}

// Starts the server based on configuration and manages graceful shutdown
func serviceAction(cobra *cobra.Command, _ []string) {
	log.Info().Msg("Starting the service...")
	serviceWatcher(getContext())
}

// serviceWatcher continuously checks the service status and performs actions
func serviceWatcher(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Info().Msg("Stopping the service...")
			return
		default:
			log.Info().Msg("Checking the service...")
			watcher.Start(ctx)
			time.Sleep(time.Duration(cfg.Service.CheckInterval) * time.Second)
		}
	}
}
