package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thomas-illiet/terrapi-runner/internal/client"
)

// Command-line flags default values
var (
	serverCmd = &cobra.Command{
		Use:   "start",
		Short: "Start workflow",
		Run:   startAction,
		Args:  cobra.NoArgs,
	}

	deploymentID      uint = 0
	managerServerAddr      = "localhost:8080"
	stateServerAddr        = "localhost:8080"
)

// Initialization of CLI flags and viper config binding
func init() {
	rootCmd.AddCommand(serverCmd)

	//region global parameters
	serverCmd.PersistentFlags().String("manager-server", managerServerAddr, "Address of the runner manager")
	viper.SetDefault("manager.address", managerServerAddr)
	_ = viper.BindPFlag("manager.address", serverCmd.PersistentFlags().Lookup("manager-server"))

	serverCmd.PersistentFlags().String("state-server", stateServerAddr, "Address of the state manager")
	viper.SetDefault("state.address", stateServerAddr)
	_ = viper.BindPFlag("state.address", serverCmd.PersistentFlags().Lookup("state-server"))
	//endregion

	//region start parameters
	serverCmd.PersistentFlags().Uint("deployment", deploymentID, "Deployment identifier")
	//endregion
}

// Starts the server based on configuration and manages graceful shutdown
func startAction(_ *cobra.Command, _ []string) {
	fmt.Println("Starting the workflow...")
	fmt.Println("Retrieving configuration...")
	webclient := client.NewClient(cfg)
	webclient.GetDeployment(deploymentID)

	fmt.Println("Cloning repository...")
	fmt.Println("Execute deployment...")
}
