package config

import (
	"os"
	"path/filepath"

	"github.com/chasinglogic/appdirs"
	"github.com/config-source/cli/client"
	"github.com/config-source/cli/utils"
	"gopkg.in/yaml.v3"
)

type Config struct {
	BaseURL string

	Token string `json:"-" yaml:"-"`
}

var Client *client.Client
var Current Config

func ConfigFile() string {
	app := appdirs.New("cdb")
	configDir := app.UserConfig()
	configFile := filepath.Join(configDir, "config.yaml")
	return configFile
}

func DefaultConfig() Config {
	return Config{
		Token:   os.Getenv("CDB_TOKEN"),
		BaseURL: os.Getenv("CDB_BASE_URL"),
	}
}

func LoadConfig() error {
	fh, err := os.Open(ConfigFile())
	if err != nil && !os.IsNotExist(err) {
		utils.Debug("unable to open config file", err)
		return err
	} else if err != nil {
		err := os.MkdirAll(appdirs.New("cdb").UserConfig(), 0700)
		if err != nil {
			utils.Debug("unable to create config directory", err)
			return err
		}

		Current = DefaultConfig()
	} else {
		err := yaml.NewDecoder(fh).Decode(&Current)
		if err != nil {
			utils.Debug("unable to decode config", err)
			return err
		}
	}

	if os.Getenv("CDB_TOKEN") != "" {
		Current.Token = os.Getenv("CDB_TOKEN")
	}

	Client = client.New(Current.Token, Current.BaseURL)
	return nil
}
