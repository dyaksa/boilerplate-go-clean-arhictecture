package bootstrap

import (
	"context"

	"github.com/dyaksa/boilerplate-go-clean-arhictecture/pkg/log"
	"github.com/dyaksa/boilerplate-go-clean-arhictecture/pkg/log/logrus"
	"github.com/dyaksa/boilerplate-go-clean-arhictecture/pqsql"
)

type Application struct {
	Env      *Env
	Log      log.Logger
	Postgres pqsql.Client
}

func App(ctx context.Context) *Application {
	app := &Application{
		Env: NewEnv(ctx),
	}

	ll, err := logrus.New(
		logrus.WithLevel("info"),
		logrus.WithJSONFormatter(),
	)

	if err != nil {
		panic(err)
	}

	app.Log = ll
	app.Postgres = NewPostgres(app.Env, app.Log)

	return app
}

func (app *Application) CloseConnection() {
	CloseConnection(app.Postgres, app.Log)
}
