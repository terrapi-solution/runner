package cmd

import (
	"github.com/spf13/cobra"
)

// Command-line flags default values
var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start workflow",
		Args:  cobra.NoArgs,
	}
)

// Initialization of CLI flags and viper config binding
func init() {
	rootCmd.AddCommand(startCmd)
}
