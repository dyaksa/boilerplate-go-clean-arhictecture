package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dyaksa/boilerplate-go-clean-arhictecture/api/route"
	"github.com/dyaksa/boilerplate-go-clean-arhictecture/bootstrap"
	"github.com/dyaksa/boilerplate-go-clean-arhictecture/pkg/log"
	"github.com/valyala/fasthttp"
)

func main() {
	ctx := context.Background()
	app := bootstrap.App(ctx)

	defer app.CloseConnection()

	env := app.Env
	router := app.Router
	l := app.Log
	db := app.Postgres
	crypto := app.Crypto

	timeout := time.Duration(env.ContextTimeout) * time.Second

	route.Setup(env, timeout, db, l, crypto, router)

	server := &fasthttp.Server{
		Handler: app.WrapHandler(app.Router.Handler),
	}

	go func() {
		l.Info(fmt.Sprintf("%s is running on port %s", env.AppName, env.Port))
		if err := server.ListenAndServe(fmt.Sprintf(":%s", env.Port)); err != nil {
			l.Error("failed to start server", log.Error("error", err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	l.Info("shutting down server")

	if err := server.Shutdown(); err != nil {
		l.Error("failed to shutdown server", log.Error("error", err))
	}

	l.Info("server stopped")

}
