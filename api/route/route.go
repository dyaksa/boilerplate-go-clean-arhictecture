package route

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func Setup(r *router.Router) {
	group := r.Group("/api")

	group.GET("/health", func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.Write([]byte("OK"))
	})
}
