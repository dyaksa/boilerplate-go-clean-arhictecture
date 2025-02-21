package bootstrap

import "context"

type Application struct {
	Env *Env
}

func App(ctx context.Context) *Application {
	app := &Application{}
	app.Env = NewEnv(ctx)
	return app
}
