package main

import (
	"github.com/spf13/viper"
)

// Function for config reading.
func readConfig() error {
	viper.SetConfigFile("config.yml")
	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	return nil
}
