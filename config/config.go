package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string `mapstructure:"server_address"`
	AdminPassword string `mapstructure:"admin_password"`
	DbUri         string `mapstructure:"db_uri"`
	InfuraURL     string `mapstructure:"infura_url"`
}

func LoadConfig(path string) Config {
	var config Config

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// err := viper.ReadConfig()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("env Read Error : &w", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("env Marshal Error : &w", err)
	}

	return config
}
