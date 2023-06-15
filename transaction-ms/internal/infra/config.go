package infra

import (
	"os"

	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

const (
	ENV_FILE = ".env"
)

type Config struct {
	ApiHost             string `env:"API_HOST,required"`
	ApiPort             string `env:"API_PORT,required"`
	DatabaseHost        string `env:"DATABASE_HOST,required"`
	DatabasePort        string `env:"DATABASE_PORT,required"`
	DatabaseName        string `env:"DATABASE_NAME,required"`
	DatabaseUser        string `env:"DATABASE_USER,required"`
	DatabasePassword    string `env:"DATABASE_PASSWORD,required"`
	AuthorizationMsHost string `env:"AUTHORIZATION_MS_HOST,required"`
	AuthorizationMsPort string `env:"AUTHORIZATION_MS_PORT,required"`
}

func NewConfig() (Config, error) {
	if err := loadEnv(); err != nil {
		return Config{}, err
	}

	config := Config{}
	if err := env.Parse(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func loadEnv() error {
	if _, err := os.Stat(ENV_FILE); os.IsNotExist(err) {
		return nil
	}

	return godotenv.Load(ENV_FILE)
}
