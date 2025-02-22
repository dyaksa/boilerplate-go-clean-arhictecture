package domain

import (
	"github.com/dyaksa/encryption-pii/crypto/types"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID       `json:"id"`
	Name      string          `json:"name"`
	Email     types.AESCipher `json:"email" full_text_search:"true"`
	EmailBidx string          `json:"email_bidx"`
	Password  string          `json:"password"`
}

type UserRepository interface{}
