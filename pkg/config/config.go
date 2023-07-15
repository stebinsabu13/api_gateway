package config

import "github.com/spf13/viper"

type Config struct {
	Port       string `mapstructure:"PORT"`
	AuthSvcUrl string `mapstructure:"AUTH_SVC_URL"`
	NoteSuvUrl string `mapstructure:"NOTE_SVC_URL"`
}

func LoadConfig() (c Config, err error) {

	viper.AddConfigPath("./")

	viper.SetConfigName(".env")

	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	if err = viper.Unmarshal(&c); err != nil {
		return
	}
	return
}
