package tokenutils

import (
	"time"

	"github.com/dyaksa/boilerplate-go-clean-arhictecture/domain"
	jwt "github.com/golang-jwt/jwt"
)

func CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	claims := &domain.JwtCustomClaims{
		ID:   user.ID,
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func CreateRefreshAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	claims := &domain.JwtCustomClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(expiry)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func ExtractIDFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidKey
		}

		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return "", jwt.ErrInvalidKey
	}

	return claims["id"].(string), nil
}
