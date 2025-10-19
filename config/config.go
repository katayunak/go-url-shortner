package config

import (
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Postgres      PostgresConfig
	Redis         RedisConfig
	MaxURLLength  int           `env:"MAX_URL_LENGTH"`
	MaxURLCache   time.Duration `env:"MAX_URL_CACHE"`
	CacheOperator string        `env:"CACHE_OPERATOR"`
}
type PostgresConfig struct {
	Host          string `env:"POSTGRES_HOST"`
	Port          string `env:"POSTGRES_PORT"`
	Username      string `env:"POSTGRES_USER"`
	Password      string `env:"POSTGRES_PASSWORD"`
	Database      string `env:"POSTGRES_DATABASE"`
	SSLMode       string `env:"POSTGRES_SSLMODE"`
	Timezone      string `env:"POSTGRES_TIMEZONE"`
	MigrationFile string `env:"POSTGRES_MIGRATION_FILE"`
}

type RedisConfig struct {
	Host       string `env:"REDIS_HOST"`
	Port       string `env:"REDIS_PORT"`
	Password   string `env:"REDIS_PASSWORD"`
	DB         int    `env:"REDIS_DB"`
	TLSEnabled bool   `env:"REDIS_TLS_ENABLED"`
}

type Configuration interface {
	Load() (*Config, error)
}

func NewConfig() Configuration {
	return &Config{}
}

func (c *Config) Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	err = env.Parse(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
