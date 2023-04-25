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
// Added some few comments in here
var Development, Staging, Production Config

func InitConfig() Config {
	env := os.Getenv("ENVIRONMENT")
	switch env {
	case "development":
		AppConfig = Development
	case "staging":
		AppConfig = Staging
	case "production":
		AppConfig = Production
	default:
		AppConfig = Production
	}

	mergo.Merge(AppConfig, DefaultConfig)

	return AppConfig
}
