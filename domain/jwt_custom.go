package domain

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type JwtCustomClaims struct {
	ID   uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
	jwt.StandardClaims
}
