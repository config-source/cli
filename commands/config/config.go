package config

import "github.com/spf13/cobra"

var Command = &cobra.Command{
	Use: "config <subcommand>",
	Aliases: []string{
		"c",
		"cfg",
	},
}

func init() {
	Command.AddCommand(getConfigCmd)
}
