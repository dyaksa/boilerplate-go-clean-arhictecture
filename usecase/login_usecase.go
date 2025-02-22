package usecase

import (
	"context"
	"time"

	"github.com/dyaksa/boilerplate-go-clean-arhictecture/domain"
	"github.com/dyaksa/boilerplate-go-clean-arhictecture/infrastructure/crypto"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	crypto         crypto.Crypto
	contextTimeout time.Duration
}

func NewLoginUsecase(ur domain.UserRepository, timeout time.Duration, crypto crypto.Crypto) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: ur,
		crypto:         crypto,
		contextTimeout: timeout,
	}
}

func (uc *loginUsecase) GetUserByEmail(ctx context.Context, email_bidx string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	email := uc.crypto.HashString(email_bidx)
	user, err := uc.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
