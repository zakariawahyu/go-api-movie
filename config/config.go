package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	AppVersion string
	Server     Server
	Postgres   Postgres
	Redis      Redis
}

type Server struct {
	Port         string
	Development  bool
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

type Redis struct {
	Host     string
	Port     string
	Password string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(`config/config.yaml`)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
