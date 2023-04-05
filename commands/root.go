package commands

import (
	"fmt"
	"os"

	"github.com/config-source/cli/client"
	"github.com/spf13/cobra"
)

var api *client.Client

var root = &cobra.Command{
	Use:   "cdb",
	Short: "Command line interface for your configuration database.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// TODO: get token and url from config
		api = client.New("", "http://localhost:3000")
	},
}

func init() {
	root.AddCommand(getCmd)
	root.AddCommand(envCmd)
}

func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
