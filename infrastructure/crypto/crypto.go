package crypto

import (
	"github.com/dyaksa/encryption-pii/crypto"
	"github.com/dyaksa/encryption-pii/crypto/core"
)

type Crypto interface {
	AESFunc() func() (core.PrimitiveAES, error)
}

type derivaleCrypto struct {
	c *crypto.Crypto
}

func (d *derivaleCrypto) AESFunc() func() (core.PrimitiveAES, error) {
	return d.c.AESFunc()
}

func New() (Crypto, error) {
	c, err := crypto.New(crypto.Aes256KeySize, crypto.WithInitHeapConnection())
	if err != nil {
		return nil, err
	}

	return &derivaleCrypto{c: c}, nil
}
