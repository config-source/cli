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
		err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load config file %s: %s", config.ConfigFile(), err)
		}
		return nil
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
