package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server   Server
	Database Database
}

type Server struct {
	Port string
}

type Database struct {
	Region    string
	TableName string
}

func LoadConfig(path string, env string) (*Config, error) {

	viperConfig := viper.New()
	viperConfig.AddConfigPath(path)
	viperConfig.SetConfigName(env)
	viperConfig.SetConfigType("yaml")

	if err := viperConfig.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	config := new(Config)

	if err := viperConfig.Unmarshal(config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return config, nil
}
