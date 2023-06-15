package pkg

import (
	"os"

	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

const (
	ENV_FILE = ".env"
)

type Config struct {
	RedisQueueHost                 string `env:"REDIS_QUEUE_HOST,required"`
	RedisQueuePort                 string `env:"REDIS_QUEUE_PORT,required"`
	RedisQueueTransactionTopicName string `env:"REDIS_QUEUE_TRANSACTION_TOPIC_NAME,required"`
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
