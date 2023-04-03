package config

import "github.com/spf13/viper"

var config = new(Config)

type Config struct {
	RunMode string
	Server  struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	MySql struct {
		Host           string `mapstructure:"host"`
		Port           string `mapstructure:"port"`
		Username       string `mapstructure:"username"`
		Password       string `mapstructure:"password"`
		Database       string `mapstructure:"database"`
		MaxIdleCon     int    `mapstructure:"max_idle_connections"`
		MaxOpenCon     int    `mapstructure:"max_open_connections"`
		MaxConLifeTime int    `mapstructure:"max_connection_lifetime"`
		LogLevel       int    `mapstructure:"log_level"`
	} `mapstructure:"mySql"`
}

func init() {
	viper.AddConfigPath("config/")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
}

func Get() Config {
	return *config
}
