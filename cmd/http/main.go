package main

import (
	"github.com/spf13/viper"
	"log"
)

const (
	configPath = "././"
	configName = "config"
	configType = "yaml"
)

func main() {

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	server := NewServer()

	server.Run()
}
