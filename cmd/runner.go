package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Command-line flags default values
var (
	runnerCmd = &cobra.Command{
		Use:   "runner",
		Short: "Start runner as fire and forget",
		Run:   runnerAction,
		Args:  cobra.NoArgs,
	}

	deploymentID uint = 0
)

// Initialization of CLI flags and viper config binding
func init() {
	startCmd.AddCommand(runnerCmd)

	//region global parameters
	runnerCmd.PersistentFlags().String("manager-server", managerServerAddr, "Address of the runner manager")
	viper.SetDefault("manager.address", managerServerAddr)
	_ = viper.BindPFlag("manager.address", runnerCmd.PersistentFlags().Lookup("manager-server"))

	runnerCmd.PersistentFlags().String("state-server", stateServerAddr, "Address of the state manager")
	viper.SetDefault("state.address", stateServerAddr)
	_ = viper.BindPFlag("state.address", runnerCmd.PersistentFlags().Lookup("state-server"))
	//endregion

	//region start parameters
	runnerCmd.PersistentFlags().Uint("deployment", deploymentID, "Deployment identifier")
	//endregion
}

// Starts the server based on configuration and manages graceful shutdown
func runnerAction(_ *cobra.Command, _ []string) {
	fmt.Println("Starting the workflow...")
	fmt.Println("Retrieving configuration...")

	fmt.Println("Cloning repository...")
	fmt.Println("Execute deployment...")
}
