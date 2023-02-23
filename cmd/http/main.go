package main

import (
	"log"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	configPath = "./"
	configName = "config"
	configType = "yaml"
)

func main() {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	production, err := zap.NewProduction()
	if err != nil {
		return
	}

	server := NewServer(production, ":8888")

	server.Run()
}
