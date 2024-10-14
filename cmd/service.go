package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Command-line flags default values
var (
	serviceCmd = &cobra.Command{
		Use:   "service",
		Short: "Start runner as service",
		Run:   serviceAction,
		Args:  cobra.NoArgs,
	}
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
}

// Starts the server based on configuration and manages graceful shutdown
func serviceAction(_ *cobra.Command, _ []string) {
	fmt.Println("Starting the workflow...")
	fmt.Println("Retrieving configuration...")

	fmt.Println("Cloning repository...")
	fmt.Println("Execute deployment...")
}
