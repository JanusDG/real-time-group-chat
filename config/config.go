package config

import (
		"fmt"
		"github.com/spf13/viper"

)

type Configurations struct {
	Server         ServerConfigurations
	IMPORTANT_VAR  int 
	DEBUG_ON       bool
}

type ServerConfigurations struct {
	Port int
}


func GetConf() *Configurations {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	
	viper.SetDefault("IMPORTANT_VAR", 42069)

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return &configuration

}