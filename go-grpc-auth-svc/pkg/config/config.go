package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port                 string `mapstructure:"PORT"`
	DBUrl                string `mapstructure:"DB_URL"`
	JWTSecretKey         string `mapstructure:"JWT_SECRET_KEY"`
	TokenExpirationHours int64  `mapstructure:"TOKEN_EXPIRATION_HOURS"`
}

func LoadConfig() (config Config, err error) {
	// viper.AddConfigPath("./pkg/config/envs")
	// viper.SetConfigName("dev")
	// viper.SetConfigType("env")

	viper.AutomaticEnv()

	viper.BindEnv("PORT")
	viper.BindEnv("DB_URL")
	viper.BindEnv("JWT_SECRET_KEY")
	viper.BindEnv("TOKEN_EXPIRATION_HOURS")

	// err = viper.ReadInConfig()
	// if err != nil {
	// return
	// }
	err = viper.Unmarshal(&config)
	return
}
