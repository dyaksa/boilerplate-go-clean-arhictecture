package bootstrap

import (
	"context"

	"github.com/dyaksa/boilerplate-go-clean-arhictecture/infrastructure/crypto"
	"github.com/dyaksa/boilerplate-go-clean-arhictecture/infrastructure/pqsql"
	"github.com/dyaksa/boilerplate-go-clean-arhictecture/pkg/log"
	"github.com/dyaksa/boilerplate-go-clean-arhictecture/pkg/log/logrus"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type Application struct {
	Env      *Env
	Router   *router.Router
	Log      log.Logger
	Postgres pqsql.Client
	Crypto   crypto.Crypto
}

func App(ctx context.Context) *Application {
	app := &Application{
		Env:    NewEnv(ctx),
		Router: router.New(),
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
	app.Crypto = NewDerivaleCrypto(app.Log)

	return app
}

func (app *Application) CloseConnection() {
	CloseConnection(app.Postgres, app.Log)
}

func (app *Application) WrapHandler(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		// next adalah handler yang akan dipanggil
		next(ctx)

		// Log informasi request, termasuk latency
		app.Log.Info("request",
			log.String("path", string(ctx.Path())),
			log.String("method", string(ctx.Method())),
			log.String("ip", ctx.RemoteIP().String()),
			log.Any("status", ctx.Response.Header.StatusCode()),
		)
	}
}
