package initializers

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost string `mapstructure:"DB_HOST"`
	DBUser string `mapstructure:"DB_USERNAME"`
	DBPort string `mapstructure:"DB_PORT"`
	DBPass string `mapstructure:"DB_PASSWORD"`
	DBName string `mapstructure:"DB_DATABASE"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
