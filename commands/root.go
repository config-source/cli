package commands

import (
	"fmt"
	"os"

	"github.com/config-source/cli/commands/config"
	"github.com/config-source/cli/commands/env"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cdb",
	Short: "Command line interface for your configuration database.",
}

func init() {
	rootCmd.AddCommand(config.Command)
	rootCmd.AddCommand(env.Command)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
