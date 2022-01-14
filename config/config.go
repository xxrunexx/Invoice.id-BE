package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost    string `yaml:"host"`
	DBPort    string `yaml:"port"`
	DBUser    string `yaml:"user"`
	DBPass    string `yaml:"pass"`
	DBName    string `yaml:"name"`
	JWTsecret string `yaml:"secret"`
	// ServerPort string `yaml:"SERVER_PORT"`
	// ServerHost string `yaml:"SERVER_HOST"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	err = viper.Unmarshal(&config)
	fmt.Println("Isi config", config)

	return config, err
}
