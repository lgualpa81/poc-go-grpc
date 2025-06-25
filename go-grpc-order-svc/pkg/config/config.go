package config

import "github.com/spf13/viper"

type Config struct {
	Port          string `mapstructure:"PORT"`
	DBUrl         string `mapstructure:"DB_URL"`
	ProductSvcUrl string `mapstructure:"PRODUCT_SVC_URL"`
}

func LoadConfig() (config Config, err error) {
	// viper.AddConfigPath("./pkg/config/envs")
	// viper.SetConfigName("dev")
	// viper.SetConfigType("env")

	viper.AutomaticEnv()

	viper.BindEnv("PORT")
	viper.BindEnv("DB_URL")
	viper.BindEnv("PRODUCT_SVC_URL")

	// err = viper.ReadInConfig()

	// if err != nil {
	// 	return
	// }

	err = viper.Unmarshal(&config)

	return
}
