package main

import (
	"github.com/spf13/viper"
)

func readConfig() error {
	viper.SetConfigFile("config.yml")
	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	return nil
}
