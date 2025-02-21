package route

import (
	"github.com/fasthttp/router"
)

func Setup(r *router.Router) {
	r.Group("/api")
}
