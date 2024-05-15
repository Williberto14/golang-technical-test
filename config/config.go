package config

import "github.com/spf13/viper"

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
