package bootstrap

import (
	"context"
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Env struct {
	AppName  string `env:"APP_NAME" default:"boilerplate-go-clean-architecture"`
	AppEnv   string `env:"APP_ENV" default:"development"`
	Port     string `env:"APP_PORT" default:"8080"`
	LogLevel string `env:"LOG_LEVEL" default:"info"`

	DBHost string `env:"DB_HOST" default:"localhost"`
	DBPort string `env:"DB_PORT" default:"5432"`
	DBUser string `env:"DB_USER" default:"postgres"`
	DBPass string `env:"DB_PASS" default:"postgres"`
	DBName string `env:"DB_NAME" default:"postgres"`
	DBSSL  string `env:"DB_SSL" default:"disable"`
}

func NewEnv(ctx context.Context) *Env {
	if err := godotenv.Load(); err != nil {
		slog.WarnContext(ctx, "failed to load env file")
	}

	env := Env{}

	if err := envconfig.Process(ctx, &env); err != nil {
		slog.WarnContext(ctx, "failed to process env vars")
	}

	return &env
}
