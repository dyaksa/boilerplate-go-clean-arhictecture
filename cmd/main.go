package main

import (
	"context"
	"fmt"

	"github.com/dyaksa/boilerplate-go-clean-arhictecture/api/route"
	"github.com/dyaksa/boilerplate-go-clean-arhictecture/bootstrap"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	ctx := context.Background()
	app := bootstrap.App(ctx)

	env := app.Env

	r := router.New()
	route.Setup(r)

	fasthttp.ListenAndServe(fmt.Sprintf(":%s", env.Port), r.Handler)
}
