package cmd

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/terrapi-solution/protocol/activity/v1"
	"github.com/terrapi-solution/runner/internal/client"
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
	runnerCmd.PersistentFlags().String("controller-server", controllerServerAddr, "Address of the controller")
	viper.SetDefault("controller.address", controllerServerAddr)
	_ = viper.BindPFlag("controller.address", runnerCmd.PersistentFlags().Lookup("controller-server"))

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
	log.Info().Msg("Starting the runner...")
	//watcher.Processing(deploymentID, getContext())

	rpcClient := client.NewClient()
	request := &activity.InsertRequest{
		Deployment: int32(1),
		Pointer:    activity.Pointer_POINTER_STDOUT,
		Message:    "scannerStdout.Text()",
	}
	rpcClient.Activity.Insert(context.Background(), request)
}
