package config

const AppName = "product:content-service"

type Config struct {
	Environment string `env:"ENVIRONMENT, default=development"`
	App         string `env:"APP, default=task-manager"`
	LogLevel    string `env:"LOG_LEVEL, default=debug"`
	Mode        string

	ServerIP string `env:"SERVER_IP, default=127.0.0.1"`
	HTTPPort string `env:"HTTP_PORT, default=:7081"`

	Postgres *DBConfig `env:",prefix=POSTGRES_"`
}

type DBConfig struct {
	Host     string `env:"HOST, default=localhost"`
	Port     uint   `env:"PORT, default=5432"`
	User     string `env:"USER, default=pguser"`
	Password string `env:"PASSWORD, default=qwerty"`
	Database string `env:"DATABASE, default=task_manager"`
}

// Logger config
type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}
