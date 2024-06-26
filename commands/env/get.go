package env

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/config-source/cli/config"
	"github.com/spf13/cobra"
)

var envGetCmd = &cobra.Command{
	Use:   "get <environment-name>",
	Short: "Get environment information by name",
	RunE: func(cmd *cobra.Command, args []string) error {
		env, err := config.Client.GetEnvironmentByName(context.Background(), args[0])
		if err != nil {
			return err
		}

		output, err := json.MarshalIndent(env, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(output))
		return nil
	},
}

func init() {
	Command.AddCommand(envGetCmd)
}
