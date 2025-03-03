package usecase

import (
	"context"
	"time"

	"github.com/dyaksa/boilerplate-go-clean-arhictecture/domain"
	"github.com/dyaksa/boilerplate-go-clean-arhictecture/infrastructure/crypto"
	"github.com/dyaksa/boilerplate-go-clean-arhictecture/pkg/tokenutils"
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

func (lu *loginUsecase) GetUserByEmail(ctx context.Context, email_bidx string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, lu.contextTimeout)
	defer cancel()

	email := lu.crypto.HashString(email_bidx)
	return lu.userRepository.GetUserByEmail(ctx, email)
}

func (lu *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutils.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutils.CreateRefreshAccessToken(user, secret, expiry)
}
