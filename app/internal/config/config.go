package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server   Server
	Database Database
}

type Server struct {
	Port           string
	RequestTimeout int
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
		return nil, fmt.Errorf("error reading config file: %w", err)
	}
	config := new(Config)

	if err := viperConfig.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return config, nil
}
