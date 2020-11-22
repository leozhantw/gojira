package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Organization string `yaml:"organization"`
}

func New() (Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}

	viper.AddConfigPath(home)
	viper.SetConfigName(".gojira")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	return cfg, err
}
