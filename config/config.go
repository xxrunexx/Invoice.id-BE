package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost    string `mapstructure:"DB_HOST"`
	DBPort    string `mapstructure:"DB_PORT"`
	DBUser    string `mapstructure:"DB_USER"`
	DBPass    string `mapstructure:"DB_PASS"`
	DBName    string `mapstructure:"DB_NAME"`
	JWTsecret string `mapstructure:"JWT_SECRET"`
	// Save for later
	// ServerPort string `mapstructure:"SERVER_PORT"`
	// ServerHost string `mapstructure:"SERVER_HOST"`
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
