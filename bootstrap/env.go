package bootstrap

import (
	"context"
	"log/slog"

	"github.com/sethvargo/go-envconfig"
)

type Env struct {
	AppName string `env:"APP_NAME" default:"boilerplate-go-clean-architecture"`
	AppEnv  string `env:"APP_ENV" default:"development"`
	Port    string `env:"PORT" default:"8080"`
}

func NewEnv(ctx context.Context) *Env {
	env := Env{}

	if err := envconfig.Process(ctx, &env); err != nil {
		slog.WarnContext(ctx, "failed to process env vars")
	}

	return &env
}
