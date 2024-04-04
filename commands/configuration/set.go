package configuration

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/config-source/cdb"
	"github.com/config-source/cli/config"
	"github.com/spf13/cobra"
)

var (
	env   string
	key   string
	value string
)

func valueAsString(cv cdb.ConfigValue) string {
	switch v := cv.Value().(type) {
	case string:
		return v
	case int:
		return fmt.Sprintf("%d", v)
	case float64, float32:
		return fmt.Sprintf("%f", v)
	case bool:
		return fmt.Sprintf("%t", v)
	default:
		return ""
	}
}

var setConfigCmd = &cobra.Command{
	Use: "set",
	RunE: func(cmd *cobra.Command, args []string) error {
		environment, err := config.Client.GetEnvironmentByNameOrID(env)
		if err != nil {
			return err
		}

		configKey, err := config.Client.GetConfigKeyByNameOrID(key)
		if err != nil {
			return err
		}

		configValue := cdb.ConfigValue{
			EnvironmentID: environment.ID,
			ConfigKeyID:   configKey.ID,
		}

		switch configKey.ValueType {
		case cdb.TypeString:
			configValue.StrValue = &value
		case cdb.TypeBoolean:
			parsed := strings.ToLower(value) == "true"
			configValue.BoolValue = &parsed
		case cdb.TypeInteger:
			parsed, err := strconv.Atoi(value)
			if err != nil {
				return err
			}

			configValue.IntValue = &parsed
		case cdb.TypeFloat:
			parsed, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}

			configValue.FloatValue = &parsed
		default:
			return errors.New("somehow couldn't find the data type of the config key")
		}

		created, err := config.Client.SetConfiguration(context.Background(), configValue)
		if err != nil {
			return err
		}

		fmt.Printf("Set %s=%s for %s\n", configKey.Name, valueAsString(created), environment.Name)
		return nil
	},
}

func init() {
	setConfigCmd.Flags().StringVarP(&env, "environment", "e", "", "The environment you want to set the value for, accepts an environment name or ID.")
	setConfigCmd.Flags().StringVarP(&key, "key", "k", "", "The configuration key you want to set the value for, accepts a key name or ID.")
	setConfigCmd.Flags().StringVarP(&value, "value", "v", "", "The value you want to set the config key to.")
	setConfigCmd.MarkFlagRequired("environment") // nolint:errcheck
	setConfigCmd.MarkFlagRequired("key")         // nolint:errcheck
	setConfigCmd.MarkFlagRequired("value")       // nolint:errcheck

	Command.AddCommand(setConfigCmd)
}
