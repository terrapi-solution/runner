package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Command-line flags default values
var (
	serverCmd = &cobra.Command{
		Use:   "start",
		Short: "Start integrated server",
		Run:   startAction,
		Args:  cobra.NoArgs,
	}

	managerServerAddr = "localhost:8085"
	stateServerAddr   = "localhost:8080"
)

// Initialization of CLI flags and viper config binding
func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.PersistentFlags().Uint("deployment", 0, "Deployment identifier")
	viper.BindPFlag("deployment.id", serverCmd.PersistentFlags().Lookup("deployment"))

	serverCmd.PersistentFlags().String("manager-server", managerServerAddr, "Address of the worker manager")
	viper.SetDefault("manager.address", managerServerAddr)
	viper.BindPFlag("manager.address", serverCmd.PersistentFlags().Lookup("manager-server"))

	serverCmd.PersistentFlags().String("state-server", stateServerAddr, "Address of the state manager")
	viper.SetDefault("state.address", stateServerAddr)
	viper.BindPFlag("state.address", serverCmd.PersistentFlags().Lookup("state-server"))
}

// Starts the server based on configuration and manages graceful shutdown
func startAction(_ *cobra.Command, _ []string) {
}
