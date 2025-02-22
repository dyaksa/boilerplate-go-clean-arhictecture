package bootstrap

import (
	"github.com/dyaksa/boilerplate-go-clean-arhictecture/infrastructure/crypto"
	"github.com/dyaksa/boilerplate-go-clean-arhictecture/pkg/log"
)

func NewDerivaleCrypto(l log.Logger) crypto.Crypto {
	c, err := crypto.New()
	if err != nil {
		l.Error("failed to create crypto", log.Error("error", err))
	}

	return c
}
