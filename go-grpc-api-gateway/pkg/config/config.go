package config

import "github.com/spf13/viper"

// we pull the data from our environment file into our API Gateway.
type Config struct {
	Port          string `mapstructure:"PORT"`
	AuthSvcUrl    string `mapstructure:"AUTH_SVC_URL"`
	ProductSvcUrl string `mapstructure:"PRODUCT_SVC_URL"`
	OrderSvcUrl   string `mapstructure:"ORDER_SVC_URL"`
}

func LoadConfig() (c Config, err error) {
	// viper.AddConfigPath("./pkg/config/envs")
	// viper.SetConfigName("dev")
	// viper.SetConfigType("env")

	viper.AutomaticEnv()

	viper.BindEnv("PORT")
	viper.BindEnv("AUTH_SVC_URL")
	viper.BindEnv("PRODUCT_SVC_URL")
	viper.BindEnv("ORDER_SVC_URL")

	// err = viper.ReadInConfig()
	// if err != nil {
	// 	return
	// }
	err = viper.Unmarshal(&c)
	return
}
