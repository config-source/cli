package commands

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var envGetCmd = &cobra.Command{
	Use:   "get <environment-name>",
	Short: "Get environment information by name",
	RunE: func(cmd *cobra.Command, args []string) error {
		env, err := api.Environments().GetByName(context.Background(), args[0])
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
