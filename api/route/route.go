package route

import (
	"time"

	"github.com/dyaksa/boilerplate-go-clean-arhictecture/bootstrap"
	"github.com/dyaksa/boilerplate-go-clean-arhictecture/infrastructure/crypto"
	"github.com/dyaksa/boilerplate-go-clean-arhictecture/infrastructure/pqsql"
	"github.com/dyaksa/boilerplate-go-clean-arhictecture/pkg/log"
	"github.com/fasthttp/router"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db pqsql.Client, l log.Logger, crypto crypto.Crypto, r *router.Router) {
	publicGroup := r.Group("/api")

	NewUserRoute(env, timeout, db, l, crypto, publicGroup)
}
