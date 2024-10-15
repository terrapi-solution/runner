package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/terrapi-solution/protocol/health"
	"github.com/terrapi-solution/runner/internal/client"
)

// Command-line flags default values
var (
	healthCmd = &cobra.Command{
		Use:   "health",
		Short: "Health",
		Long:  "Gets the health of the specified service",
		Run:   healthAction,
		Args:  cobra.NoArgs,
	}

	// Default service name to check
	defaultServiceName = "controller"
)

// Initialization of CLI flags and viper config binding
func init() {
	rootCmd.AddCommand(healthCmd)

	healthCmd.Flags().StringVar(&defaultServiceName, "service", "controller", "A provided int")
}

// Starts the server based on configuration and manages graceful shutdown
func healthAction(cobra *cobra.Command, _ []string) {
	c := client.NewClient()

	result, err := c.Health.Check(getContext(), &health.CheckRequest{Service: defaultServiceName})

	fmt.Println(result.Status)
	fmt.Println(err)
}
