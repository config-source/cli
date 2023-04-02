package commands

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use: "get <environment-name> <configuration-key-name>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		env, key := args[0], args[1]

		value, err := api.Environments().GetConfigValue(ctx, env, key)
		if err == nil {
			fmt.Println(value.Value())
		}

		return err
	},
}
