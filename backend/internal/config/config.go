package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database `mapstructure:"mssql"`
	Server   `mapstructure:"server"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type Server struct {
	Host string     `mapstructure:"shost"`
	Port int        `mapstructure:"sport"`
	Echo EchoServer `mapstructure:"echo"`
}

type EchoServer struct {
	Debug                         bool
	EnableCORSMiddleware          bool
	EnableLoggerMiddleware        bool
	EnableRecoverMiddleware       bool
	EnableRequestIDMiddleware     bool
	EnableTrailingSlashMiddleware bool
	EnableSecureMiddleware        bool
	EnableCacheControlMiddleware  bool
}

func Load(path string) (config Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(err)
	}
	return
}
