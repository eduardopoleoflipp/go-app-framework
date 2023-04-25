package config

import (
	"os"

	"github.com/imdario/mergo"
)

type Config struct {
	Name          string
	MyConfig      string
	DefaultConfig string
	NotSetConfig  string
}

var DefaultConfig = Config{DefaultConfig: "defaultValue"}
var AppConfig Config

// need to be overwritten by the actual app
var development, staging, production Config

func InitConfig() Config {
	env := os.Getenv("ENVIRONMENT")
	switch env {
	case "development":
		AppConfig = development
	case "staging":
		AppConfig = staging
	case "production":
		AppConfig = production
	default:
		AppConfig = production
	}

	mergo.Merge(AppConfig, DefaultConfig)

	return AppConfig
}
