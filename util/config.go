package util

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	 DBDriver string `mapstructure:"DB_DRIVER"`
	 DBSource string `mapstructure:"DB_SOURCE"`
	 ServerAdress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found, falling back to environment variables")
		} else {
			return
		}
	}

	err = viper.Unmarshal(&config)
	return
}