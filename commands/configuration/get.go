package configuration

import (
	"cmp"
	"context"
	"fmt"
	"slices"

	"github.com/config-source/cdb"
	"github.com/config-source/cli/config"
	"github.com/config-source/cli/table"
	"github.com/spf13/cobra"
)

func valueToRow(cv cdb.ConfigValue) []string {
	repr := ""
	switch cv.ValueType {
	case cdb.TypeString:
		repr = *cv.StrValue
	case cdb.TypeInteger:
		repr = fmt.Sprintf("%d", *cv.IntValue)
	case cdb.TypeFloat:
		repr = fmt.Sprintf("%f", *cv.FloatValue)
	case cdb.TypeBoolean:
		repr = fmt.Sprintf("%t", *cv.BoolValue)
	default:
		repr = "UNKNOWN VALUE!"
	}

	return []string{
		cv.Name,
		repr,
		fmt.Sprintf("%t", cv.Inherited),
	}
}

func printConfigTable(values []cdb.ConfigValue) {
	tbl := table.Table{
		Headings: []string{"Key", "Value", "Inherited"},
		Rows:     make([][]string, len(values)),
	}

	slices.SortFunc(values, func(a, b cdb.ConfigValue) int {
		return cmp.Compare(a.Name, b.Name)
	})

	for idx, value := range values {
		tbl.Rows[idx] = valueToRow(value)
	}

	fmt.Println(tbl)
}

var getConfigCmd = &cobra.Command{
	Use: "get <environment-name> [configuration-key-name]",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		var env, key string
		env = args[0]
		if len(args) > 1 {
			key = args[1]
		}

		if key != "" {
			value, err := config.Client.GetConfigurationValue(ctx, env, key)
			if err != nil {
				return err
			}

			fmt.Println(value.Value())
		} else {
			values, err := config.Client.GetConfiguration(ctx, env)
			if err != nil {
				return err
			}

			printConfigTable(values)
		}

		return nil
	},
}
