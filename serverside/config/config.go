package config

import (
	"log"

	"github.com/spf13/viper"
)

type Configurations struct {
	Server        ServerConfigurations
	IMPORTANT_VAR int
	DEBUG_ON      bool
}

type ServerConfigurations struct {
	Port     int
	DEBUG_ON bool
}

func GetConf() *Configurations {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	viper.SetDefault("IMPORTANT_VAR", 42069)

	var configuration Configurations

	if err := viper.Unmarshal(&configuration); err != nil {
		log.Printf("Unable to decode into struct, %v", err)
	}

	return &configuration

}
