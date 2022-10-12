package config

import "github.com/spf13/viper"

func ParseConfig() (*viper.Viper, error) {
	v := viper.New()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	v.SetConfigFile("./config/config.yaml")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v, nil
}
