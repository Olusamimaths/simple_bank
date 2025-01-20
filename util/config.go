package util

import (
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER" validate:"required"`
	DBSource            string        `mapstructure:"DB_SOURCE" validate:"required"`
	ServerAdress        string        `mapstructure:"SERVER_ADDRESS" validate:"required"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMETRIC_KEY" validate:"required"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION" validate:"required"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION" validate:"required"`
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
	if err != nil {
		log.Printf("Unable to unmarshal config: %v", err)
		return
	}

	validate := validator.New()
	err = validate.Struct(config)
	if err != nil {
		log.Printf("config validation error: %v", err)
		return
	}

	return
}
