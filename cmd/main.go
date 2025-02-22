package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dyaksa/boilerplate-go-clean-arhictecture/api/route"
	"github.com/dyaksa/boilerplate-go-clean-arhictecture/bootstrap"
	"github.com/dyaksa/boilerplate-go-clean-arhictecture/pkg/log"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	ctx := context.Background()
	app := bootstrap.App(ctx)

	defer app.CloseConnection()

	env := app.Env

	r := router.New()
	route.Setup(r)

	server := &fasthttp.Server{
		Handler: r.Handler,
	}

	go func() {
		app.Log.Info(fmt.Sprintf("%s is running on port %s", env.AppName, env.Port))
		if err := server.ListenAndServe(fmt.Sprintf(":%s", env.Port)); err != nil {
			app.Log.Error("failed to start server", log.Error("error", err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	app.Log.Info("shutting down server")

	if err := server.Shutdown(); err != nil {
		app.Log.Error("failed to shutdown server", log.Error("error", err))
	}

	app.Log.Info("server stopped")

}
