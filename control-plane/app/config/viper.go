package config

import (
	"strings"

	"github.com/spf13/viper"
)

func LoadConfiguration(environment Environment) error {
	if environment == Production {
		viper.SetConfigName("config.prod")
	} else {
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
