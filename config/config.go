package config

import (
	"github.com/stakkato95/service-engineering-go-lib/config"
)

type Config struct {
	ServerPort    string `mapstructure:"SERVER_PORT"`
	UsersService  string `mapstructure:"USERS_SERVICE"`
	UsersGrpcPort string `mapstructure:"USERS_GRPC_PORT"`
	TweetsService string `mapstructure:"TWEETS_SERVICE"`
}

var AppConfig Config

func init() {
	config.Init(&AppConfig, Config{})
}

func UsersGrpc() string {
	return AppConfig.UsersService + AppConfig.UsersGrpcPort
}
