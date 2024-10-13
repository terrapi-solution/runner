package cmd

import (
	"github.com/spf13/cobra"
	"github.com/terrapi-solution/runner/internal/config"
)

var (
	rootCmd = &cobra.Command{
		Use:   "terrapi-runner",
		Short: "Terrapi runner",

		SilenceErrors: false,
		SilenceUsage:  true,

		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	cfg *config.Config
)

func init() {
	cfg = config.Load()
	cobra.OnInitialize(setupConfig)

	rootCmd.PersistentFlags().BoolP("help", "h", false, "Show the help, so what you see now")
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Print the current version of that tool")
}

// Run parses the cmd line arguments and executes the program.
func Run() error {
	return rootCmd.Execute()
}
