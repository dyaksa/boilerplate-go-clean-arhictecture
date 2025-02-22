package domain

import (
	"context"

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

func (u User) ScanDestinations() []interface{} {
	return []interface{}{
		&u.ID,
		&u.Name,
		&u.Email,
	}
}

func (u *User) To() any {
	return u
}

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email_bidx string) (*User, error)
}
