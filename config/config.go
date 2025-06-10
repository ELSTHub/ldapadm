package config

import (
	"github.com/spf13/viper"
	"os"
)

func ConfInit() error {
	viper.SetConfigName("ldapadm")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/ldapadm")
	home, _ := os.UserHomeDir()
	viper.AddConfigPath(home + "/.etc")
	viper.AddConfigPath("./etc")
	viper.AddConfigPath("../etc")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
