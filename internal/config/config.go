package config

import (
	"fmt"
	"time"
)

const AppName = "product:content-service"

type Config struct {
	Environment string `env:"ENVIRONMENT, default=development"`
	App         string `env:"APP, default=content_service"`
	AppVersion  string `env:"APP, default=1.0.0"`

	ServerIP string `env:"SERVER_IP, default=127.0.0.1"`
	HTTPPort string `env:"HTTP_PORT, default=:8080"`

	Postgres *DBConfig `env:",prefix=POSTGRES_"`
	Logger   *Logger   `env:",prefix=LOGGER_"`
}

type DBConfig struct {
	Host     string `env:"HOST, default=localhost"`
	Port     uint   `env:"PORT, default=5555"`
	User     string `env:"USER, default=admin"`
	Password string `env:"PASSWORD, default=147ajt369"`
	Database string `env:"DATABASE, default=dbo_content_service"`
	SSLMode  string `env:"SSLMODE, default=disable"`

	MaxIdleConnections    int32         `env:"MAX_IDLE_CONNECTIONS,default=25"`
	MaxOpenConnections    int32         `env:"MAX_OPEN_CONNECTIONS,default=25"`
	ConnectionMaxLifetime time.Duration `env:"CONNECTION_MAX_LIFETIME,default=5m"`
}

// Logger config
type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	LogLevel          string `env:"LEVEL, default=debug"`
}

func (c *Config) PostgresURL() string {
	return c.Postgres.PostgresURL()
}

func (c *DBConfig) PostgresURL() string {
	if c.User == "" {
		return fmt.Sprintf(
			"host=%s port=%d  dbname=%s sslmode=%s",
			c.Host,
			c.Port,
			c.Database,
			c.SSLMode,
		)
	}

	if c.Password == "" {
		return fmt.Sprintf(
			"host=%s port=%d user=%s  dbname=%s sslmode=%s",
			c.Host,
			c.Port,
			c.User,
			c.Database,
			c.SSLMode,
		)
	}

	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.Database,
		c.SSLMode,
	)
}
