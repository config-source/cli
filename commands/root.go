package commands

import (
	"fmt"
	"os"

	"github.com/config-source/cli/commands/configuration"
	"github.com/config-source/cli/commands/env"
	"github.com/config-source/cli/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cdb",
	Short: "Command line interface for your configuration database.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return config.LoadConfig()
	},
}

func init() {
	rootCmd.AddCommand(configuration.Command)
	rootCmd.AddCommand(env.Command)
	rootCmd.AddCommand(setupCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
