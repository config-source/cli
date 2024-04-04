package config

import (
	"os"
	"path/filepath"

	"github.com/chasinglogic/appdirs"
	"gopkg.in/yaml.v3"
)

type Config struct {
	BaseURL string

	Token string `json:"-" yaml:"-"`
}

var Current Config

func ConfigFile() string {
	app := appdirs.New("cdb")
	configDir := app.UserConfig()
	configFile := filepath.Join(configDir, "config.yaml")
	return configFile
}

func LoadConfig() error {
	fh, err := os.Open(ConfigFile())
	if err != nil && !os.IsNotExist(err) {
		return err
	} else if err != nil {
		err := os.MkdirAll(appdirs.New("cdb").UserConfig(), 0700)
		if err != nil {
			return err
		}
	} else {
		err := yaml.NewDecoder(fh).Decode(&Current)
		if err != nil {
			return err
		}
	}

	Current.Token = os.Getenv("CDB_TOKEN")
	return nil
}
