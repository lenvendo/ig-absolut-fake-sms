package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	"github.com/lenvendo/ig-absolut-fake-sms/lib/db"
	"github.com/lenvendo/ig-absolut-fake-sms/lib/log"
)

var cfg Config

func Get() Config {
	return cfg
}

func InitConfig() Config {
	cfg, err := LoadConfig("APP")
	if err != nil {
		fmt.Printf("Ошибка загрузки конфига, %v\n", err)
		os.Exit(1)
	}
	if cfg.AppName == "" {
		cfg.AppName = "app"
	}

	return cfg
}

type Config struct {
	AppName string `envconfig:"NAME"`

	HTTP  *HTTPServerConfig
	Debug bool
	ENV   string

	Database *db.PostgreSQLConfig `envconfig:"DATABASE"`

	Nats *NatsConfig `envconfig:"NATS"`

	IsDebug bool `envconfig:"IS_DEBUG"`
	Env     string

	Logger log.Config
}

type NatsConfig struct {
	Host       string `envconfig:"HOST"`
	Port       int    `envconfig:"PORT"`
	RetryLimit int    `envconfig:"RETRYLIMIT"`
	WaitLimit  int    `envconfig:"WAITLIMIT"`
	Username   string `envconfig:"USERNAME"`
	Password   string `envconfig:"PASSWORD"`
}

type HTTPServerConfig struct {
	Port string
}

// LoadConfig подтягивает все env переменные из файлов .env, .env.local
// и создает новый Config с помощью envconfig
func LoadConfig(envconfigPrefix string) (Config, error) {
	cfg := &Config{}
	_ = godotenv.Load(".env")
	_ = godotenv.Overload(".env.local")

	err := envconfig.Process(envconfigPrefix, cfg)
	if err != nil {
		return *cfg, err
	}

	return *cfg, nil
}

func IsDevOrTestEnv(config *Config) bool {
	return config.ENV == "test" || config.ENV == "dev"
}
