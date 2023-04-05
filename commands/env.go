package commands

import "github.com/spf13/cobra"

var envCmd = &cobra.Command{
	Use: "environment <subcommand>",
	Aliases: []string{
		"e",
		"env",
	},
}

func init() {
	envCmd.AddCommand(envGetCmd)
	envCmd.AddCommand(envTreeCmd)
}
