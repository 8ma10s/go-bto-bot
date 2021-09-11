package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	AppName       string  `mapstructure:"app_name"`
	Discord       Discord `mapstructure:"discord"`
	CommandPrefix string  `mapstructure:"command_prefix"`
}

type Discord struct {
	Token string `mapstructure:"token"`
}

func Load() (*Config, error) {
	setConfigDirectory()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	viper.SetEnvPrefix("BTO")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var cfg Config

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func setConfigDirectory() {
	const FILENAME = "config"
	const EXTENSION = "yaml"
	const CONFIGPATH = "configs/"

	viper.SetConfigName(FILENAME)
	viper.SetConfigType(EXTENSION)
	viper.AddConfigPath(CONFIGPATH)
}
