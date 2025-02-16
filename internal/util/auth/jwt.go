package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/nghiatk54/go_ecommerce_api/global"
)

// payload claims
type PayloadClaims struct {
	jwt.RegisteredClaims
}

// generate token
func GenerateTokenJwt(payload jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(global.Config.Jwt.API_SECRET))
}

// create token
func CreateToken(uuidToken string) (string, error) {
	// 1. set time expiration
	timeEx := global.Config.Jwt.JWT_EXPIRATION
	if timeEx == "" {
		timeEx = "1h"
	}
	expirationTime, err := time.ParseDuration(timeEx)
	if err != nil {
		return "", err
	}
	now := time.Now()
	expiresAt := now.Add(expirationTime)
	// 2. generate token
	return GenerateTokenJwt(&PayloadClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    "shopdevgo",
			Subject:   uuidToken,
		},
	})
}

// parse token
func ParseJwtTokenSubject(token string) (*jwt.RegisteredClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.Jwt.API_SECRET), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*jwt.RegisteredClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// verify token
func VerifyTokenSubject(token string) (*jwt.RegisteredClaims, error) {
	claims, err := ParseJwtTokenSubject(token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
